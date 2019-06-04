package insight

import (
	"fmt"
	"github.com/dikhan/insight_goclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

const labelResourceName = "insight_label"
const labelResourceId = "acceptance_label"

var labelResourceStateId = fmt.Sprintf("%s.%s", labelResourceName, labelResourceId)

var testLabelCreateConfig string

var createLabelName = "My New Awesome Label"
var createLabelColor = "ff0000"

func init() {

	configTemplate := `
   provider insight {
     api_key = "%s"
     region  = "%s"
   }

   resource %s %s {
     name = "%s"
     color = "%s"
   }`

	testLabelCreateConfig = fmt.Sprintf(configTemplate, apiKey, region, labelResourceName, labelResourceId, createLabelName, createLabelColor)
}

type checkExists func(client insight_goclient.InsightClient, id string) error

func labelExists() checkExists {
	return func(client insight_goclient.InsightClient, id string) error {
		_, err := client.GetLabel(id)
		return err
	}
}

func checkDestroy(resourceStateId string, checkExists checkExists) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		leClient := testAccProvider.Meta().(insight_goclient.InsightClient)
		if len(s.Modules) != 0 {
			if s.Modules[0].Resources[resourceStateId] != nil {
				id := s.Modules[0].Resources[resourceStateId].Primary.ID
				if err := checkExists(leClient, id); err != nil {
					if !strings.Contains(err.Error(), "404 Not Found") {
						return fmt.Errorf("received an error retrieving resource %s - %s", id, err.Error())
					}
				} else {
					return fmt.Errorf(fmt.Sprintf("Resource %s still exists remotely", id))
				}
			}
		} else {
			return fmt.Errorf(fmt.Sprintf("an error occurred while processing the resource %s, state file does not have enough information", resourceStateId))
		}
		return nil
	}
}

func TestAccInsightLabel_Create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkDestroy(labelResourceStateId, labelExists()),
		Steps: []resource.TestStep{
			{
				Config: testLabelCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(labelResourceStateId, "name", createLabelName),
					resource.TestCheckResourceAttr(labelResourceStateId, "color", createLabelColor),
				),
			},
		},
	})
}
