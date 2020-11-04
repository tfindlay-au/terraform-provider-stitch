# terraform-provider-stitch

### Objective
This code base provides link between Hashicorp's Terraform tools and the Stitch Data API.

### Using
Before you begin you will need to create an account at StitchData (https://www.stitchdata.com/)

Once you have an account you will need to purchase an enterprise plan to access Stitch Connect (API Access)
https://www.stitchdata.com/docs/developers/stitch-connect

You will then be able to login and goto your account settings and under `API access keys` generate a key to be used
with this plugin in terraform.

*NOTE*: Without "enterprise" access you will only have access to very limited Import API found here:
https://www.stitchdata.com/docs/developers

Using the plugin with in your terraform scripts should look like:

```hcl-terraform
terraform {
  required_providers {
    stitch = {
      source  = "stitchdata.com/provider/stitch"
      version = "~> 0.1.0"
    }
  }
}

provider "stitch" {
  api_key = "<My API Key>"
}

resource "stitch_source" "source_database" {
  display_name = "test_database"
  type = "<sourceType>"
  properties = {
    host     = "myHost",
    port     = 1234,
    user     = "myUsername",
    password = "myPassword"
  }
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
To build from the source, run:
```shell
go build -o terraform-provider-stitch_v0.1.0
mv terraform-provider-stitch_v0.1.0 ~/.terraform.d/plugins/stitchdata.com/provider/stitch/0.1.0/darwin_amd64/terraform-provider-stitch_v0.1.0
rm -Rf ~/.terraform.d/plugin-cache/stitchdata.com
```

This will produce a binary file that can be placed in your terraform plugins folder. eg `~/.terraform.s/plugins/`

### Testing
```shell
go build -o terraform-provider-stitch_v0.1.0
mv terraform-provider-stitch_v0.1.0 ~/.terraform.d/plugins/stitchdata.com/provider/stitch/0.1.0/darwin_amd64/terraform-provider-stitch_v0.1.0
rm -Rf ~/.terraform.d/plugin-cache/stitchdata.com
cd terraform
rm -Rf .terraform terraform.tfstate terraform.tfstate.backup
terraform init
terraform apply --auto-approve
cd ..
```

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

`TODO` Create makefile
