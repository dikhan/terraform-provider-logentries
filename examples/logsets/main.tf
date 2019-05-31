# Specify the provider and access details
provider insight {
  api_key = "${var.api_key}"
  region  = "eu"
}

resource insight_logsets my_logset {
  name        = "My Log Set"
  description = "Description about my log set"
  logs_info   = []
  user_data   = {}
}
