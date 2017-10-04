# Specify the provider and access details
provider "logentries" {
  api_key = "${var.api_key}"
}

resource "logentries_logsets" "my_logset" {
  name = "My log Set"
  description = "some description goes here"
}

resource "logentries_logs" "my_log" {
  name = "My super log"
  source_type = "token"
  token_seed = ""
  structures = []
  logsets_info = ["${logentries_logsets.my_logset.id}"]
  user_data = {
    "le_agent_filename" = "/var/log/anaconda.log",
    "le_agent_follow" = "true"
  }
}