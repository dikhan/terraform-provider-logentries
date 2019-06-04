package insight

import (
	"fmt"
	"github.com/dikhan/insight_goclient"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

const logsResourceName = "insight_log"
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
	provider "insight" {
	  api_key = "%s"
	  region  = "%s"
	}

	resource insight_logset acceptance_logset {
	  name = "Acceptance Log Set for log acc tests"
	  description = "some description goes here"
	}

	resource "%s" "%s" {
	  	name = "%s"
	  	source_type = "token"
	  	token_seed = ""
	  	structures = []
	  	logsets_info = ["${insight_logset.acceptance_logset.id}"]
	  	user_data {
	  		agent_filename = "%s"
			agent_follow   = "%s"
	  	}
	}`

	testLogCreateConfig = fmt.Sprintf(configTemplate, apiKey, region, logsResourceName, logsResourceId, createLogName, createLogAgentFileName, createLogAgentFollow)
	testLogUpdatedConfig = fmt.Sprintf(configTemplate, apiKey, region, logsResourceName, logsResourceId, updatedLogName, updatedLogAgentFileName, updatedLogAgentFollow)
}

func logExists() checkExists {
	return func(client insight_goclient.InsightClient, id string) error {
		_, err := client.GetLog(id)
		return err
	}
}

func TestAccInsightLog_Create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkDestroy(logResourceStateId, logExists()),
		Steps: []resource.TestStep{
			{
				Config: testLogCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logResourceStateId, "name", createLogName),
					resource.TestCheckResourceAttr(logResourceStateId, "source_type", "token"),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.%", "2"),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.agent_filename", createLogAgentFileName),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.agent_follow", createLogAgentFollow),
				),
			},
		},
	})
}

func TestAccInsightLog_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkDestroy(logResourceStateId, logExists()),
		Steps: []resource.TestStep{
			{
				Config: testLogCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logResourceStateId, "name", createLogName),
					resource.TestCheckResourceAttr(logResourceStateId, "source_type", "token"),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.%", "2"),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.agent_filename", createLogAgentFileName),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.agent_follow", createLogAgentFollow),
				),
			},
			{
				Config: testLogUpdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logResourceStateId, "name", updatedLogName),
					resource.TestCheckResourceAttr(logResourceStateId, "source_type", "token"),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.%", "2"),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.agent_filename", updatedLogAgentFileName),
					resource.TestCheckResourceAttr(logResourceStateId, "user_data.agent_follow", updatedLogAgentFollow),
				),
			},
		},
	})
}
