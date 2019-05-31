# Specify the provider and access details
provider insight {
  api_key = "${var.api_key}"
  region  = "eu"
}

resource insight_logsets my_logset {
  name        = "My log Set"
  description = "some description goes here"
}

resource insight_logs my_log {
  name         = "My super log"
  source_type  = "token"
  token_seed   = ""
  structures   = []
  logsets_info = ["${insight_logsets.my_logset.id}"]

  user_data = {
    "le_agent_filename" = "/var/log/anaconda.log"
    "le_agent_follow"   = "true"
  }
}
