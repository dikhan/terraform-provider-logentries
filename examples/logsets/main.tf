# Specify the provider and access details
provider "logentries" {
  api_key = "${var.api_key}"
}

resource "logentries_logsets" "my_logset" {
  name = "My Log Set"
  description = "Description about my log set"
  logs_info = []
  user_data = {}
}