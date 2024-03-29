package stitch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"time"
)

// PURPOSE:
// The Destination object represents a destination. Destinations are the data warehouses into which Stitch writes data
// The Destination Type object contains the information needed to configure a destination.
// Schema Reference here: https://www.stitchdata.com/docs/developers/stitch-connect/api#destinations
var destinationSchema = map[string]*schema.Schema{
	"destination_id": {
		Type:        schema.TypeInt,
		Description: "A unique identifier for this destination.",
		Computed:    true,
		Required:    false,
	},
	"name": {
		Type:        schema.TypeString,
		Description: "The name for the destination.",
		Computed:    false,
		Required:    true,
	},
	"display_name": {
		Type:        schema.TypeString,
		Description: "The display name of the destination.",
		Computed:    false,
		Required:    true,
	},
	"created_at": {
		Type:        schema.TypeInt,
		Description: "The time at which the destination object was created.",
		Computed:    true,
		Required:    false,
	},
	"deleted_at": {
		Type:        schema.TypeInt,
		Description: "The time at which the destination object was deleted.",
		Computed:    true,
		Required:    false,
	},
	"paused_at": {
		Type:        schema.TypeString,
		Description: "The name for the destination.",
		Computed:    true,
		Required:    false,
	},
	"stitch_client_id": {
		Type:        schema.TypeInt,
		Description: "The ID of the Stitch client account associated with the destination.",
		Computed:    true,
		Required:    false,
	},
	"system_paused_at": {
		Type:        schema.TypeInt,
		Description: "If the connection was paused by the system, the time the pause began. Otherwise, or if the connection is active, this will be null.",
		Computed:    true,
		Required:    false,
	},
	"updated_at": {
		Type:        schema.TypeInt,
		Description: "The time at which the destination object was last updated.\n",
		Computed:    true,
		Required:    false,
	},
	"type": {
		Type:        schema.TypeInt,
		Description: "The destination type. Must be one of: redshift, s3, postgres, snowflake etc",
		Computed:    true,
		Required:    false,
	},
}

func destination() *schema.Resource {
	return &schema.Resource{
		Description:   "Destinations are the data warehouses into which Stitch writes data.",
		CreateContext: destinationCreate,
		ReadContext:   destinationList,
		UpdateContext: destinationUpdate,
		DeleteContext: destinationDelete,
		Schema:        destinationSchema,
		SchemaVersion: 4,
	}
}

// Create POST /v4/destinations
func destinationCreate(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Container for feedback
	var d diag.Diagnostics

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v4/destinations", c.HostURL), nil)
	if err != nil {
		return diag.FromErr(err)
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

	return d
}

// Update PUT /v4/destinations/{destination_id}
func destinationUpdate(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics
	d = append(d, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Not implemented.",
		Detail:   "The function destinationUpdate() in the stitch provider plugin has not been implemented",
	})
	return d
}

// List GET /v4/destinations
func destinationList(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics
	d = append(d, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Not implemented.",
		Detail:   "The function destinationUpdate() in the stitch provider plugin has not been implemented",
	})
	return d
}

// Delete DELETE /v4/destinations/{destination_id}
func destinationDelete(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v4/destinations/{destination_id}", c.HostURL), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = c.doRequest(req)
	if err != nil {
		return diag.FromErr(err)
	}

	return d
}
