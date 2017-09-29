# Specify the provider and access details
provider "logentries" {
  api_key = "${var.api_key}"
}

resource "logentries_tags" "my_tag" {
  name = "My App Failures"
  type = "Alert"
  patterns = ["[error]"]
  sources = ["${var.source_log_id}"]
  labels = [
    {
      name = "my-app-failures"
      color = "ff0000"
      reserved = false
      sn = 1881
    }
  ]
  actions = [
    {
      type = "Alert"
      enabled = true
      min_matches_count = 1
      min_matches_period = "Hour"
      min_report_count = 1
      min_report_period = "Hour"
      targets = [
        {
          type = "Pagerduty"
          alert_content_set = {
            le_context = "true"
            le_log_link = "true"
          }
          params_set {
            description = "Log Error",
            service_key = "${var.pagerduty_key}"
          }
        }
      ]
    }
  ]
}