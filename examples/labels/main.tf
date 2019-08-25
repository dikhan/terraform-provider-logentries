# Specify the provider and access details
provider insight {
  api_key = "${var.api_key}"
  region  = "eu"
}

resource insight_label my_label {
  name  = "My Label"
  color = "ff0000"
}
