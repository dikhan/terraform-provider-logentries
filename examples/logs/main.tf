# Specify the provider and access details
provider insight {
  api_key = "${var.api_key}"
  region  = "eu"
}

resource insight_logset my_logset {
  name        = "My log Set"
  description = "some description goes here"
}

resource insight_log my_log {
  name           = "My super log"
  source_type    = "token"
  logset_ids     = ["${insight_logset.my_logset.id}"]
  agent_filename = "/var/log/anaconda.log"
  agent_follow   = "true"
}
