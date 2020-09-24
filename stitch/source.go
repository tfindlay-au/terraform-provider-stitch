package stitch

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// PURPOSE:
// The Source Type object contains the information needed to configure a data source.
// An object representing a data source. Sources are the databases, APIs, and other data applications that
// Stitch replicates data from. Sources can be retrieved in a list or individually by ID.

func source() *schema.Resource {
	return &schema.Resource{
		ReadContext: sourceGet,
		Schema: map[string]*schema.Schema{
			"coffees": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// List of types (eg platform.mysql, platform.facebook) GET /v4/source-types
// Get type details (eg platform.jira) GET /v4/source-types/{source_type}

// Create POST /v4/sources
// Update PUT /v4/sources/{source_id}
// 		Pause PUT /v4/sources/{source_id}
//		UnPause PUT /v4/sources/{source_id}
// List all GET /v4/sources

// Get Details (specific) GET /v4/sources/{source_id}
func sourceGet(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics
	// Do stuff in here
	return d
}

// Delete DELETE /v4/sources/{source_id}

// Create Access Token for Source POST /v4/sources/{source_id}/tokens
// Get Access Token for Source GET /v4/sources/{source_id}/tokens
// Delete Access Token for Source DELETE /v4/sources/{source_id}/tokens/{token_id}
