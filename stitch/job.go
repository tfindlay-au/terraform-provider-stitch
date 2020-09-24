package stitch

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// PURPOSE:
// The three-step process by which Stitch replicates data.
// A replication job includes three distinct steps: Extraction, preparation, and loading

// Job Object:
// {
// "job_name": "116078.120643.sync.c12fb0a7-7e4a-11e9-abdc-0edc2c318fba"
// }
var jobSchema = map[string]*schema.Schema{
	"job_name": {
		Type:     schema.TypeString,
		Computed: false,
		Required: true,
	},
}

func job() *schema.Resource {
	return &schema.Resource{
		CreateContext: startJob,
		DeleteContext: stopJob,
		Schema:        jobSchema,
	}
}

// Start a job POST /v4/sources/{source_id}/sync
func startJob(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics
	// Do stuff in here...
	return d
}

// Stop a job DELETE /v4/sources/{source_id}/sync
func stopJob(ctx context.Context, r *schema.ResourceData, m interface{}) diag.Diagnostics {
	var d diag.Diagnostics
	// Do stuff in here...
	return d
}
