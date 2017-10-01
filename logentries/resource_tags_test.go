package logentries

import (
	"testing"
	resource "github.com/hashicorp/terraform/helper/resource"
	"fmt"
	"os"
	"github.com/dikhan/logentries_goclient"
)


const tagsResourceName = "logentries_tags"
const tagsResourceId = "acceptance_tag"

var sourceId = os.Getenv("SOURCE_ID")
var labelId = os.Getenv("LABEL_ID")

var tagsResourceStateId = fmt.Sprintf("%s.%s", tagsResourceName, tagsResourceId)

var testTagsCreateConfig string
var testTagsUpdatedConfig string

var createTagName = "My App Failures"
var createTagPatterns = "[error]"
var createTagActionParamSetDescription = "Log Error"

var updatedTagName = "My Update App Failures"
var updatedTagPatterns = "[error] AND failure and 502"
var updatedTagActionParamSetDescription = "Log Error description updated"

func init() {

	configTemplate := `
	provider "logentries" {
	  api_key = "%s"
	}

	resource "%s" "%s" {
	  name = "%s"
	  type = "Alert"
	  patterns = ["%s"]
	  sources = ["%s"]
	  labels = [
		{
		  id = "%s"
		  name = "logentries_client_label_test"
		  color = "ff0000"
		  reserved = false
		  sn = 2004
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
				description = "%s",
				service_key = "MY_PD_KEY"
			  }
			}
		  ]
		}
	  ]
	}`

	testTagsCreateConfig = fmt.Sprintf(configTemplate, os.Getenv("LOGENTRIES_API_KEY"), tagsResourceName, tagsResourceId, createTagName, createTagPatterns, sourceId, labelId, createTagActionParamSetDescription)
	testTagsUpdatedConfig = fmt.Sprintf(configTemplate, os.Getenv("LOGENTRIES_API_KEY"), tagsResourceName, tagsResourceId, updatedTagName, updatedTagPatterns, sourceId, labelId, updatedTagActionParamSetDescription)
}

func tagExists() checkExists {
	return func(leClient logentries_goclient.LogEntriesClient, id string) error {
		_, err := leClient.Tags.GetTag(id)
		fmt.Println(err)
		if err != nil {
			return err
		}
		return nil
	}
}

func TestAccLogentriesTags_Create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t)},
		Providers:testAccProviders,
		CheckDestroy: checkDestroy(tagsResourceStateId, tagExists()),
		Steps: []resource.TestStep {
			resource.TestStep{
				Config: testTagsCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tagsResourceStateId, "name", createTagName),
					resource.TestCheckResourceAttr(tagsResourceStateId, "type", "Alert"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.0", createTagPatterns),
					resource.TestCheckResourceAttr(tagsResourceStateId, "sources.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "sources.0", sourceId),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.id", labelId),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.name", "logentries_client_label_test"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.color", "ff0000"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.reserved", "false"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.sn", "2004"),

					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.enabled", "true"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.min_matches_count", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.min_matches_period", "Hour"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.min_report_count", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.min_report_period", "Hour"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.alert_content_set.%", "2"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.alert_content_set.le_context", "true"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.alert_content_set.le_log_link", "true"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.params_set.%", "2"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.params_set.description", createTagActionParamSetDescription),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.params_set.service_key", "MY_PD_KEY"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.type", "Pagerduty"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.type", "Alert"),
				),
			},
		},
	})
}


func TestAccLogentriesTags_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t)},
		Providers:testAccProviders,
		CheckDestroy: checkDestroy(tagsResourceStateId, tagExists()),
		Steps: []resource.TestStep {
			resource.TestStep{
				Config: testTagsCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tagsResourceStateId, "name", createTagName),
					resource.TestCheckResourceAttr(tagsResourceStateId, "type", "Alert"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.0", createTagPatterns),
					resource.TestCheckResourceAttr(tagsResourceStateId, "sources.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "sources.0", sourceId),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.id", labelId),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.name", "logentries_client_label_test"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.color", "ff0000"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.reserved", "false"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.sn", "2004"),

					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.enabled", "true"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.min_matches_count", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.min_matches_period", "Hour"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.min_report_count", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.min_report_period", "Hour"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.alert_content_set.%", "2"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.alert_content_set.le_context", "true"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.alert_content_set.le_log_link", "true"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.params_set.%", "2"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.params_set.description", createTagActionParamSetDescription),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.params_set.service_key", "MY_PD_KEY"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.type", "Pagerduty"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.type", "Alert"),
				),
			},
			resource.TestStep{
				Config: testTagsUpdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tagsResourceStateId, "name", updatedTagName),
					resource.TestCheckResourceAttr(tagsResourceStateId, "type", "Alert"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.0", updatedTagPatterns),
					resource.TestCheckResourceAttr(tagsResourceStateId, "sources.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "sources.0", sourceId),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.id", labelId),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.name", "logentries_client_label_test"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.color", "ff0000"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.reserved", "false"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "labels.0.sn", "2004"),

					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.enabled", "true"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.min_matches_count", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.min_matches_period", "Hour"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.min_report_count", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.min_report_period", "Hour"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.alert_content_set.%", "2"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.alert_content_set.le_context", "true"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.alert_content_set.le_log_link", "true"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.params_set.%", "2"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.params_set.description", updatedTagActionParamSetDescription),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.params_set.service_key", "MY_PD_KEY"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.targets.0.type", "Pagerduty"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "actions.0.type", "Alert"),
				),
			},
		},
	})
}