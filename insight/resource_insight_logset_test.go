package insight

import (
	"fmt"
	"github.com/Tweddle-SE-Team/insight_goclient"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

const logSetsResourceName = "insight_logset"
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
	provider insight {
	  api_key = "%s"
	  region  = "%s"
	}

	resource %s %s {
	  name = "%s"
	  description = "%s"
	  logs_info = []
	  user_data = {}
	}`

	testLogSetsCreateConfig = fmt.Sprintf(configTemplate, apiKey, region, logSetsResourceName, logSetsResourceId, createLogSetName, createLogSetDescription)
	testLogSetsUpdatedConfig = fmt.Sprintf(configTemplate, apiKey, region, logSetsResourceName, logSetsResourceId, updatedLogSetName, updatedLogSetDescription)
}

func logSetExists() checkExists {
	return func(client insight_goclient.InsightClient, id string) error {
		_, err := client.GetLogset(id)
		return err
	}
}

func TestAccLogentriesLogSets_Create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkDestroy(logSetsResourceStateId, logSetExists()),
		Steps: []resource.TestStep{
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkDestroy(logSetsResourceStateId, logSetExists()),
		Steps: []resource.TestStep{
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
