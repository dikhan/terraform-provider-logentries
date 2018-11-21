package logentries

import (
	"fmt"
	"github.com/dikhan/logentries_goclient"
	"github.com/hashicorp/terraform/helper/resource"
	"os"
	"testing"
)

const tagsResourceName = "logentries_tags"
const tagsResourceId = "acceptance_tag"

var tagsResourceStateId = fmt.Sprintf("%s.%s", tagsResourceName, tagsResourceId)

var testTagsCreateConfig string
var testTagsUpdatedConfig string

var createTagName = "My App Failures"
var createTagPatterns = "[error]"
var createTagActionParamSetDescription = "Log Error"

var updatedTagName = "My Update App Failures"
var updatedTagPatterns = "[error] AND failure and 502"
var updatedTagActionParamSetDescription = "Log Error description updated"

var sourceId = os.Getenv("SOURCE_ID")

func init() {

	configTemplate := `
	provider "logentries" {
	  api_key = "%s"
	}

	resource "logentries_logsets" "acceptance_logset" {
	  name = "Sample LogSet for tag acc tests"
	  description = "some description goes here"
	}

	resource "logentries_logs" "acceptance_log" {
	  name = "Sample Log for tag acc tests"
	  source_type = "token"
	  token_seed = ""
	  structures = []
	  logsets_info = ["${logentries_logsets.acceptance_logset.id}"]
	  user_data = {
	    "le_agent_filename" = "",
		"le_agent_follow" = "false"
	  }
	}

	resource "%s" "%s" {
	  name = "%s"
	  type = "Alert"
	  patterns = ["%s"]
	  sources = ["${logentries_logs.acceptance_log.id}"]
	  labels = []
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

	testTagsCreateConfig = fmt.Sprintf(configTemplate, apiKey, tagsResourceName, tagsResourceId, createTagName, createTagPatterns, createTagActionParamSetDescription)
	testTagsUpdatedConfig = fmt.Sprintf(configTemplate, apiKey, tagsResourceName, tagsResourceId, updatedTagName, updatedTagPatterns, updatedTagActionParamSetDescription)
}

func tagExists() checkExists {
	return func(leClient logentries_goclient.LogEntriesClient, id string) error {
		_, err := leClient.Tags.GetTag(id)
		return err
	}
}

func TestAccLogentriesTags_Create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkDestroy(tagsResourceStateId, tagExists()),
		Steps: []resource.TestStep{
			{
				Config: testTagsCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tagsResourceStateId, "name", createTagName),
					resource.TestCheckResourceAttr(tagsResourceStateId, "type", "Alert"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.0", createTagPatterns),
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkDestroy(tagsResourceStateId, tagExists()),
		Steps: []resource.TestStep{
			{
				Config: testTagsCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tagsResourceStateId, "name", createTagName),
					resource.TestCheckResourceAttr(tagsResourceStateId, "type", "Alert"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.0", createTagPatterns),
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
			{
				Config: testTagsUpdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tagsResourceStateId, "name", updatedTagName),
					resource.TestCheckResourceAttr(tagsResourceStateId, "type", "Alert"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.0", updatedTagPatterns),
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
