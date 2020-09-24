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
var destinationSchema = map[string]*schema.Schema{
	"id": {
		Type:        schema.TypeInt,
		Description: "A unique identifier for this destination.",
		Computed:    true,
		Required:    true,
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
		CreateContext: destinationCreate,
		DeleteContext: destinationDelete,
		UpdateContext: destinationUpdate,
		ReadContext:   destinationList,
		Schema:        destinationSchema,
	}
}

// List of types (eg s3, snowflake) GET /v4/destination-types
// Get type details (eg Redshift) GET /v4/destination-types/{destination_type}

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

	authString := fmt.Sprintf("Bearer %s", c.Token)
	req.Header.Add("Authorization", authString)
	req.Header.Add("Content-Type", "application/json")
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
	// Do stuff in here...
	return d
}

// List GET /v4/destinations
func destinationList(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics
	// Do stuff in here...
	return d
}

// Delete DELETE /v4/destinations/{destination_id}
func destinationDelete(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics
	// Do stuff in here...
	return d
}
