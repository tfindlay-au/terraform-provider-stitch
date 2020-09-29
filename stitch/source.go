package stitch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"time"
)

// PURPOSE:
// The Source Type object contains the information needed to configure a data source.
// An object representing a data source. Sources are the databases, APIs, and other data applications that
// Stitch replicates data from. Sources can be retrieved in a list or individually by ID.
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
		Type:        schema.TypeInt,
		Description: "The time at which the source object was created.",
		Computed:    true,
		Required:    false,
	},
	"updated_at": {
		Type:        schema.TypeInt,
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
		Type:        schema.TypeInt,
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
		Type:        schema.TypeString,
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

// List of types (eg platform.mysql, platform.facebook) GET /v4/source-types
// Get type details (eg platform.jira) GET /v4/source-types/{source_type}

// Create POST /v4/sources
func sourceCreate(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

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
	//defer response.Body.Close()

	payload := make([]map[string]interface{}, 0)
	err = json.Unmarshal(response, &payload)
	if err != nil {
		return diag.FromErr(err)
	}
	// Do stuff in here
	return d
}

// Update PUT /v4/sources/{source_id}
// 		Pause PUT /v4/sources/{source_id}
//		UnPause PUT /v4/sources/{source_id}
func sourceUpdate(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics
	// Do stuff in here
	return d
}

// List all GET /v4/sources
func sourceGetList(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics
	// Do stuff in here
	return d
}

// Get Details (specific) GET /v4/sources/{source_id}
func sourceGetDetails(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics

	// Do stuff in here
	return d
}

// Delete DELETE /v4/sources/{source_id}
func sourceDelete(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics
	// Do stuff in here
	return d
}

// Create Access Token for Source POST /v4/sources/{source_id}/tokens
// Get Access Token for Source GET /v4/sources/{source_id}/tokens
// Delete Access Token for Source DELETE /v4/sources/{source_id}/tokens/{token_id}
