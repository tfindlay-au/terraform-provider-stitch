# terraform-provider-stitch

### Objective
This code base provides link between Hashicorp's Terraform tools and the Stitch Data API.

### Using
Using the plugin with in your terraform scripts should look like:

```hcl-terraform
provider "stitch" {
  apiKey = "<My API Key>"
}


resource "stitch_source" "source_database" {
  name                        = "test_database"
  display_name                = "Sample Source via Terraform"
}

resource "stitch_destinations" "target_database" {
  name                        = "test_database"
  comment                     = "Sample Target via Terraform"
}

resource "stitch_job" "transfer_job" {
  name                        = "test_database"
  comment                     = "Sample Storage via Terraform"
  data_retention_time_in_days = 3
}

```

### Building
To build from source, run:
```shell
go build -o terraform-provider-stitch
```

This will produce a binary file that can be placed in your terraform plugins folder. eg `~/.terraform.s/plugins/`

### Installing

#### Manually

To install manually, see the guide at:
You will need to create the following directory structure:
    HOSTNAME/NAMESPACE/TYPE/VERSION/TARGET
Where:

* HOSTNAME
* NAMESPACE
* TYPE
* TARGET is the operating system like `darwin_amd64`, `linux_arm`, `windows_amd64`
* VERSION is the version of the stitch plugin eg. 0.1

For example, consider the following path to install on OSx (Mac)
```shell
~/.terraform.d/plugins/stitchdata.com/stitchdata/stitch/0.1/darwin_amd64/terraform-provider-stitch
```

#### Automatically

// TODO Create makefile
