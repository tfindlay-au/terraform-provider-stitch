package stitch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
)

// PURPOSE:
// The Source Type object contains the information needed to configure a data source.
// An object representing a data source. Sources are the databases, APIs, and other data applications that
// Stitch replicates data from. Sources can be retrieved in a list or individually by ID.
// Schema Reference: https://www.stitchdata.com/docs/developers/stitch-connect/api#source--object
var sourceSchema = map[string]*schema.Schema{
	"source_id": {
		Type:        schema.TypeInt,
		Description: "A unique identifier for this destination.",
		Computed:    true,
		Required:    false,
	},
	"properties": {
		Type:        schema.TypeMap,
		Description: "Parameters for connecting to the source, excluding any sensitive credentials. The parameters must adhere to the type of source.",
		Computed:    false,
		Required:    true,
	},
	"created_at": {
		Type:        schema.TypeString,
		Description: "The time at which the source object was created.",
		Computed:    true,
		Required:    false,
	},
	"updated_at": {
		Type:        schema.TypeString,
		Description: "The time at which the object was last updated.",
		Computed:    true,
		Required:    false,
	},
	"deleted_at": {
		Type:        schema.TypeInt,
		Description: "The time at which the source object was deleted.",
		Computed:    true,
		Required:    false,
	},
	"schedule": {
		Type:        schema.TypeMap,
		Description: "An object describing the replication schedule for the source.",
		Computed:    false,
		Optional:    true,
	},
	"check_job_name": {
		Type:        schema.TypeString,
		Description: "The name of the last connection check job that ran for the source.",
		Computed:    true,
		Required:    false,
	},
	"name": {
		Type:        schema.TypeString,
		Description: "The name of the source connection, dynamically generated from display_name. The name corresponds to the destination schema name that the data from this source will be loaded into.\n\nNames must:\n\n    Contain only lowercase alphanumerics and underscores\n    Be unique within each Stitch client account\n",
		Computed:    true,
		Required:    false,
	},
	"type": {
		Type:        schema.TypeString,
		Description: "The source type.",
		Computed:    false,
		Required:    true,
	},
	"system_paused_at": {
		Type:        schema.TypeString,
		Description: "If the connection was paused by the system, the time the pause began. Otherwise, or if the connection is active, this will be null.",
		Computed:    true,
		Required:    false,
	},
	"stitch_client_id": {
		Type:        schema.TypeInt,
		Description: "The ID of the Stitch client account associated with the source.",
		Computed:    true,
		Required:    false,
	},
	"paused_at": {
		Type:        schema.TypeInt,
		Description: "If the connection was paused by the user, the time the pause began. Otherwise, or if the connection is active, this will be `null.",
		Computed:    true,
		Required:    false,
	},
	"display_name": {
		Type:        schema.TypeString,
		Description: "The display name of the source connection.",
		Computed:    false,
		Required:    true,
	},
	"report_card": {
		Type:        schema.TypeMap,
		Description: "A description of the source’s configuration state.",
		Computed:    true,
		Required:    false,
	},
}

func source() *schema.Resource {
	return &schema.Resource{
		Description:   "Sources are the databases, APIs, and other data applications that Stitch replicates data from.",
		CreateContext: sourceCreate,
		ReadContext:   sourceGetDetails,
		UpdateContext: sourceUpdate,
		DeleteContext: sourceDelete,
		Schema:        sourceSchema,
		SchemaVersion: 4,
	}
}

// Create POST /v4/sources
func sourceCreate(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics

	c := m.(*Client)

	url := fmt.Sprintf("%s/v4/sources", c.HostURL)

	// TODO map this to the schema structure
	configuration := map[string]interface{}{
		"display_name": r.Get("display_name").(string),
		"type":         r.Get("type").(string),
		"properties": map[string]interface{}{
			"host":     r.Get("properties.host").(string),
			"port":     r.Get("properties.port").(string),
			"user":     r.Get("properties.user").(string),
			"password": r.Get("properties.password").(string),
		},
	}
	body, _ := json.Marshal(configuration)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		// Provide useful visibility if it fails to form a request object
		d = append(d, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Message Body",
			Detail:   string(body),
		})
		d = append(d, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Target URL",
			Detail:   url,
		})
		d = append(d, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to execute request",
			Detail:   err.Error(),
		})
		return d
	}

	response, err := c.doRequest(req)
	if err != nil {
		return diag.FromErr(err)
	}

	payload := make(map[string]interface{}, 1)
	err = json.Unmarshal(response, &payload)
	if err != nil {
		d = append(d, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Failed to parse JSON response.",
			Detail:   string(response),
		})
		d = append(d, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to process JSON response",
			Detail:   err.Error(),
		})
		return d
	}

	uniqueId := fmt.Sprintf("%.0f", payload["id"])
	r.SetId(uniqueId)

	return d
}

// Update PUT /v4/sources/{source_id}
// 		Pause PUT /v4/sources/{source_id}
//		UnPause PUT /v4/sources/{source_id}
func sourceUpdate(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics
	d = append(d, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Not implemented.",
		Detail:   "The function sourceUpdate() in the stitch provider plugin has not been implemented",
	})
	return d
}

// Get Details (specific) GET /v4/sources/{source_id}
func sourceGetDetails(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics

	d = append(d, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Not implemented.",
		Detail:   "The function sourceGetDetails() in the stitch provider plugin has not been implemented",
	})

	return d
}

// Delete a source /v4/sources/{source_id}
func sourceDelete(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics
	c := m.(*Client)

	url := fmt.Sprintf("%s/v4/sources/%s", c.HostURL, r.Id())

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		d = append(d, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Target URL",
			Detail:   url,
		})
		d = append(d, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to execute request",
			Detail:   err.Error(),
		})
		return d
	}

	_, err = c.doRequest(req)
	if err != nil {
		return diag.FromErr(err)
	}

	return d
}

// List all GET /v4/sources
// 		Lists the sources for an account, including active, paused, and deleted sources

// For Sources that use access tokens, implement:
// Create Access Token for Source POST /v4/sources/{source_id}/tokens
// Get Access Token for Source GET /v4/sources/{source_id}/tokens
// Delete Access Token for Source DELETE /v4/sources/{source_id}/tokens/{token_id}

// To map attributes to source types, use these API's eg:
//	{
//	"type": "platform.mysql",
//	"current_step": 1,
//	"current_step_type": "form",
//	"steps": [
//	{
//	"type": "form",
//	"properties": [
//	{
//	"name": "allow_non_auto_increment_pks",
//	"is_required": false,
//	"is_credential": false,
//	"system_provided": false,
//	"property_type": "user_provided",
//	"json_schema": {
//	"type": "string",
//	"pattern": "^(true|false)$"
//	},
//	"provided": false,
//	"tap_mutable": false
//	},
//	{
//	"name": "anchor_time",
//	"is_required": false,
//	"is_credential": false,
//	"system_provided": false,
//	"property_type": "user_provided",
// List of types (eg platform.mysql, platform.facebook) GET /v4/source-types
// Get type details (eg platform.jira) GET /v4/source-types/{source_type}
