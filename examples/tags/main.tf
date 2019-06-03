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

  target_ids = [
    "${insight_target.pagerduty.id}",
  ]
}

resource insight_target pagerduty {
  name = "pgdt"
  pagerduty_service_key = "asdasdasdasd"
}

resource insight_tag my_tag {
  name       = "My App Failures"
  type       = "Alert"
  patterns   = ["[error]"]
  source_ids = ["5a1288ab-561a-4f93-1111-6a38c6d8TEST"]
  label_ids  = ["${data.insight_label.label.label_id}"]

  action_ids = [
    "${insight_action.alert1.id}",
  ]
}
