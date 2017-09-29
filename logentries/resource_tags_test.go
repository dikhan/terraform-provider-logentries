package logentries

import (
	"testing"
	"github.com/hashicorp/terraform/helper/resource"
	"fmt"
	"os"
)


const tagsResourceName = "logentries_tags"
const tagsResourceId = "acceptance_tag"

var tagsResourceStateId = fmt.Sprintf("%s.%s", tagsResourceName, tagsResourceId)

var testTagsCreateConfig string

func init() {
	testTagsCreateConfig = fmt.Sprintf(`
	provider "logentries" {
	  api_key = "%s"
	}

	resource "logentries_tags" "acceptance_tag" {
	  name = "My App Failures"
	  type = "Alert"
	  patterns = ["[error]"]
	  sources = ["%s"]
	  labels = [
		{
		  id = "%s"
		  name = "le-terraform-provider-label"
		  color = "ff0000"
		  reserved = false
		  sn = 2005
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
				service_key = "%s"
			  }
			}
		  ]
		}
	  ]
	}`, os.Getenv("LOGENTRIES_API_KEY"), os.Getenv("SOURCE_ID"), os.Getenv("LABEL_ID"), "PD_API_KEY")
}



func TestAccLogentriesTags_Create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t)},
		Providers:testAccProviders,
		Steps: []resource.TestStep {
			resource.TestStep{
				Config: testTagsCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tagsResourceStateId, "name", ""),
				),
			},
		},
	})
}
