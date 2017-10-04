package logentries

import (
	"testing"
	"github.com/hashicorp/terraform/helper/resource"
	"fmt"
	"github.com/dikhan/logentries_goclient"
)

const logSetsResourceName = "logentries_logsets"
const logSetsResourceId = "acceptance_logset"

var logSetsResourceStateId = fmt.Sprintf("%s.%s", logSetsResourceName, logSetsResourceId)

var testLogSetsCreateConfig string
var testLogSetsUpdatedConfig string

var createLogSetName = "Acceptance Log Set"
var createLogSetDescription = "some description for the logset"

var updatedLogSetName = "My Update App Failures"
var updatedLogSetDescription = "updated description for the logset"

func init() {

	configTemplate := `
	provider "logentries" {
	  api_key = "%s"
	}

	resource "%s" "%s" {
	  name = "%s"
	  description = "%s"
	  logs_info = []
	  user_data = {}
	}`

	testLogSetsCreateConfig = fmt.Sprintf(configTemplate, apiKey, logSetsResourceName, logSetsResourceId, createLogSetName, createLogSetDescription)
	testLogSetsUpdatedConfig = fmt.Sprintf(configTemplate, apiKey, logSetsResourceName, logSetsResourceId, updatedLogSetName, updatedLogSetDescription)
}

func logSetExists() checkExists {
	return func(leClient logentries_goclient.LogEntriesClient, id string) error {
		_, _, err := leClient.LogSets.GetLogSet(id)
		return err
	}
}

func TestAccLogentriesLogSets_Create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t)},
		Providers:testAccProviders,
		CheckDestroy: checkDestroy(logSetsResourceStateId, logSetExists()),
		Steps: []resource.TestStep {
			{
				Config: testLogSetsCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logSetsResourceStateId, "name", createLogSetName),
					resource.TestCheckResourceAttr(logSetsResourceStateId, "description", createLogSetDescription),
				),
			},
		},
	})
}

func TestAccLogentriesLogSets_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t)},
		Providers:testAccProviders,
		CheckDestroy: checkDestroy(logSetsResourceStateId, logSetExists()),
		Steps: []resource.TestStep {
			{
				Config: testLogSetsCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logSetsResourceStateId, "name", createLogSetName),
					resource.TestCheckResourceAttr(logSetsResourceStateId, "description", createLogSetDescription),
				),
			},
			{
				Config: testLogSetsUpdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logSetsResourceStateId, "name", updatedLogSetName),
					resource.TestCheckResourceAttr(logSetsResourceStateId, "description", updatedLogSetDescription),
				),
			},
		},
	})
}