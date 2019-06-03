# Specify the provider and access details
provider insight {
  api_key = "${var.api_key}"
  region  = "eu"
}

data insight_label label {
  name  = "Critical"
  color = "e0e000"
}

resource insight_action alert1 {
  type               = "Alert"
  enabled            = true
  min_matches_count  = 1
  min_matches_period = "Hour"
  min_report_count   = 1
  min_report_period  = "Hour"
  targets = [
    "${insight_target.pagerduty.id}",
  ]
}

resource insight_target pagerduty {
  pagerduty_key = "asdasdasdasd"
}

resource insight_tags my_tag {
  name     = "My App Failures"
  type     = "Alert"
  patterns = ["[error]"]
  sources  = ["5a1288ab-561a-4f93-1111-6a38c6d8TEST"]
  labels   = ["${data.insight_label.label.label_id}"]

  actions = [
    "${insight_action.alert1.id}",
    "${insight_action.alert2.id}",
  ]
}
