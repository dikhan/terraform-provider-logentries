# Specify the provider and access details
provider "logentries" {
  api_key = "${var.api_key}"
}

data "logentries_labels" "label" {
  name  = "Critical"
  color = "e0e000"
}

resource "logentries_tags" "my_tag" {
  name     = "My App Failures"
  type     = "Alert"
  patterns = ["[error]"]
  sources  = ["5a1288ab-561a-4f93-1111-6a38c6d8TEST"]
  labels   = ["${data.logentries_labels.label.label_id}"]

  actions = [
    {
      type               = "Alert"
      enabled            = true
      min_matches_count  = 1
      min_matches_period = "Hour"
      min_report_count   = 1
      min_report_period  = "Hour"

      targets = [
        {
          type = "Pagerduty"

          alert_content_set = {
            le_context  = "true"
            le_log_link = "true"
          }

          params_set {
            description = "Log Error"
            service_key = "${var.pagerduty_key}"
          }
        },
      ]
    },
  ]
}
