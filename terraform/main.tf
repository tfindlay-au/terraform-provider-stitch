terraform {
  required_providers {
    stitch = {
      source  = "stitchdata.com/provider/stitch"
      version = "~> 0.1.0"
    }
  }
}

provider "stitch" {
  api_key = "<API Key goes here>"
}

resource "stitch_source" "source_database" {
  display_name = "test_database"
  type         = "platform.mysql"
  properties = {
    host     = "myHost"
    port     = "1234"
    user     = "myUsername"
    password = "myPassword"
  }
}