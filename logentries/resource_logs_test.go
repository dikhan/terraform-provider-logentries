package logentries

import (
	"testing"
	"github.com/hashicorp/terraform/helper/resource"
	"fmt"
	"github.com/dikhan/logentries_goclient"
)

const logsResourceName = "logentries_logs"
const logsResourceId = "acceptance_log"

var logResourceStateId = fmt.Sprintf("%s.%s", logsResourceName, logsResourceId)

var testLogCreateConfig string
var testLogUpdatedConfig string

var createLogName = "My New Awesome Log"
var createLogAgentFileName = "/var/log/anaconda.log"
var createLogAgentFollow = "false"

var updatedLogName = "My Updated Awesome Log"
var updatedLogAgentFileName = "/var/log/snake.log"
var updatedLogAgentFollow = "true"

func init() {

	configTemplate := `
	provider "logentries" {
	  api_key = "%s"
	}

	resource "logentries_logsets" "acceptance_logset" {
	  name = "Acceptance Log Set for log acc tests"
	  description = "some description goes here"
	}

	resource "%s" "%s" {
	  name = "%s"
	  source_type = "token"
	  token_seed = ""
	  structures = []
	  logsets_info = ["${logentries_logsets.acceptance_logset.id}"]
	  user_data = {
	    "le_agent_filename" = "%s",
		"le_agent_follow" = "%s"
	  }
	}`

	testLogCreateConfig = fmt.Sprintf(configTemplate, apiKey, logsResourceName, logsResourceId, createLogName, createLogAgentFileName, createLogAgentFollow)
	testLogUpdatedConfig = fmt.Sprintf(configTemplate, apiKey, logsResourceName, logsResourceId, updatedLogName, updatedLogAgentFileName, updatedLogAgentFollow)
}

func logExists() checkExists {
	return func(leClient logentries_goclient.LogEntriesClient, id string) error {
		_, _, err := leClient.Logs.GetLog(id)
		fmt.Println(err)
		return err
	}
}

func TestAccLogentriesLog_Create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t)},
		Providers:testAccProviders,
		CheckDestroy: checkDestroy(logResourceStateId, logExists()),
		Steps: []resource.TestStep {
			{
				Config: testLogCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logResourceStateId, "name", createLogName),
					resource.TestCheckResourceAttr(logResourceStateId, "source_type", "token"),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.%", "2"),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.le_agent_filename", createLogAgentFileName),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.le_agent_follow", createLogAgentFollow),
				),
			},
		},
	})
}

func TestAccLogentriesLog_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t)},
		Providers:testAccProviders,
		CheckDestroy: checkDestroy(logResourceStateId, logExists()),
		Steps: []resource.TestStep {
			{
				Config: testLogCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logResourceStateId, "name", createLogName),
					resource.TestCheckResourceAttr(logResourceStateId, "source_type", "token"),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.%", "2"),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.le_agent_filename", createLogAgentFileName),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.le_agent_follow", createLogAgentFollow),
				),
			},
			{
				Config: testLogUpdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logResourceStateId, "name", updatedLogName),
					resource.TestCheckResourceAttr(logResourceStateId, "source_type", "token"),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.%", "2"),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.le_agent_filename", updatedLogAgentFileName),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.le_agent_follow", updatedLogAgentFollow),
				),
			},
		},
	})
}