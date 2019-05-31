package insight

import (
	"fmt"
	"github.com/dikhan/insight_goclient"
	"github.com/hashicorp/terraform/helper/resource"
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
   provider "insight" {
     api_key = "%s"
   }

   resource "%s" "%s" {
     name = "%s"
     color = "%s"
   }`

	testLabelCreateConfig = fmt.Sprintf(configTemplate, apiKey, labelResourceName, labelResourceId, createLabelName, createLabelColor)
}

func labelExists() checkExists {
	return func(leClient insight_goclient.InsightClient, id string) error {
		_, err := leClient.Labels.GetLabel(id)
		return err
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
