package stitchApi

import "encoding/json"

type Platform struct {
	DataType        string  `json:"type"`
	CurrentStep     int     `json:"current_step"`
	CurrentStepType string  `json:"current_step_type"`
	Steps           []Step  `json:"steps"`
	Details         Details `json:"details"`
}

type Step struct {
	StepType   string           `json:"type"`
	Properties []StepProperties `json:"properties"`
}

type StepProperties struct {
	Name           string         `json:"name"`
	IsRequired     bool           `json:"is_required"`
	IsCredential   bool           `json:"is_credential"`
	SystemProvided bool           `json:"system_provided"`
	PropertyType   string         `json:"property_type"`
	Provided       bool           `json:"provided"`
	TapMutable     bool           `json:"tap_mutable"`
	JsonSchema     PropertySchema `json:"json_schema"`
}

type PropertySchema struct {
	PropertyType string `json:"type"`
	Pattern      string `json:"pattern"`
}

type Details struct {
	PricingTier                string `json:"pricing_tier"`
	PipelineState              string `json:"pipeline_state"`
	DefaultStartDate           string `json:"default_start_date"`
	DefaultSchedulingIntervale int    `json:"default_scheduling_interval"`
	Protocol                   string `json:"protocol"`
	Access                     bool   `json:"access"`
}

func (p Platform) SourceTypeValidator(r string) bool {
	// Given the requirements from stitch and a definition
	// in Terraform, match the 2 files to see if we can continue

	// For each step in the definition ...
	for i := range p.Steps {

		// For each Property ...
		for _, v := range p.Steps[i].Properties {

			// If its required by StitchData
			if v.IsRequired {
				propertyExists(&r, &v.Name)
			}
		}
	}
	return false
}

func propertyExists(input *string, fieldName *string) bool {
	// We do not know what structure the input string will be as
	// it comes from the terraform configuration file
	var parsedInput map[string]interface{}

	err := json.Unmarshal([]byte(*input), &parsedInput)
	if err != nil {
		return false
	}

	// For each element int to top of the tree...
	for k, v := range parsedInput {
		if k == *fieldName {
			return true
		}

		if k == "properties" {
			b := v.(map[string]interface{})
			for _, c := range b["properties"].([]interface{}) {

				// Cast the interface{} -> map[string]interface
				d := c.(map[string]interface{})
				if _, ok := d[*fieldName]; ok {

					return true
				}
			}
		}
	}
	return false
}
