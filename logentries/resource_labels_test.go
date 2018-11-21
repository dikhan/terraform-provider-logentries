package logentries

import (
	"fmt"
	"github.com/dikhan/logentries_goclient"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

const labelsResourceName = "logentries_labels"
const labelsResourceId = "acceptance_label"

var labelResourceStateId = fmt.Sprintf("%s.%s", labelsResourceName, labelsResourceId)

var testLabelCreateConfig string

var createLabelName = "My New Awesome Label"
var createLabelColor = "ff0000"

func init() {

	configTemplate := `
   provider "logentries" {
     api_key = "%s"
   }

   resource "%s" "%s" {
     name = "%s"
     color = "%s"
   }`

	testLabelCreateConfig = fmt.Sprintf(configTemplate, apiKey, labelsResourceName, labelsResourceId, createLabelName, createLabelColor)
}

func labelExists() checkExists {
	return func(leClient logentries_goclient.LogEntriesClient, id string) error {
		_, err := leClient.Labels.GetLabel(id)
		return err
	}
}

func TestAccLogentriesLabel_Create(t *testing.T) {
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
