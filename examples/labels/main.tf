# Specify the provider and access details
provider logentries {
  api_key = "${var.api_key}"
}

resource logentries_labels my_label {
  name  = "My Label"
  color = "ff0000"
}
