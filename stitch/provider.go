package stitch

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Parameters available for a provider definition
var providerSchema = map[string]*schema.Schema{
	"api_key": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		DefaultFunc: schema.EnvDefaultFunc("STITCH_APIKEY", nil),
	},
}

// Define the provider here
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: providerSchema,
		ResourcesMap: map[string]*schema.Resource{
			"stitch_destination": destination(),
			"stitch_source":      source(),
			"stitch_job":         job(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := rd.Get("api_key").(string)

	// Warning or errors can be collected in a slice type
	var d diag.Diagnostics

	if apiKey != "" {
		c, err := NewSession(nil, &apiKey)
		if err != nil {
			d = append(d, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Stitch client",
				Detail:   err.Error(),
			})
			return nil, d
		}

		// Success, log something useful anyway.
		d = append(d, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Stitch Session Token from " + c.HostURL,
			Detail:   c.Token,
		})

		return c, d
	} else {
		d = append(d, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to find API Key",
			Detail:   "Value:" + apiKey,
		})
		return nil, d
	}
}
