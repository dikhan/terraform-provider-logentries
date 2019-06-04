package insight

import (
	"fmt"
	"github.com/dikhan/insight_goclient"
	"github.com/hashicorp/terraform/helper/resource"
	"os"
	"testing"
)

const tagsResourceName = "insight_tag"
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
	provider insight {
	  api_key = "%s"
	  region  = "%s"
	}

	resource insight_logset acceptance_logset {
	  name = "Sample LogSet for tag acc tests"
	  description = "some description goes here"
	}

	resource insight_log acceptance_log {
	  name = "Sample Log for tag acc tests"
	  source_type = "token"
	  token_seed = ""
	  structures = []
	  logsets_info = ["${insight_logset.acceptance_logset.id}"]
	  user_data = {
	   	agent_filename = ""
			agent_follow = false
	  }
	}

	resource %s %s {
	  name = "%s"
	  type = "Alert"
	  patterns = ["%s"]
	  source_ids = ["${insight_log.acceptance_log.id}"]
	  label_ids = []
	  action_ids = []
	}`

	testTagsCreateConfig = fmt.Sprintf(configTemplate, apiKey, region, tagsResourceName, tagsResourceId, createTagName, createTagPatterns, createTagActionParamSetDescription)
	testTagsUpdatedConfig = fmt.Sprintf(configTemplate, apiKey, region, tagsResourceName, tagsResourceId, updatedTagName, updatedTagPatterns, updatedTagActionParamSetDescription)
}

func tagExists() checkExists {
	return func(client insight_goclient.InsightClient, id string) error {
		_, err := client.GetTag(id)
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
				),
			},
			{
				Config: testTagsUpdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tagsResourceStateId, "name", updatedTagName),
					resource.TestCheckResourceAttr(tagsResourceStateId, "type", "Alert"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.0", updatedTagPatterns),
				),
			},
		},
	})
}
