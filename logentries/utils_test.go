package logentries

import (
	"github.com/hashicorp/terraform/terraform"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/dikhan/logentries_goclient"
	"strings"
	"fmt"
)

type checkExists func(leClient logentries_goclient.LogEntriesClient, id string) error

func checkDestroy(resourceStateId string, checkExists checkExists) resource.TestCheckFunc {
	return func (s *terraform.State) error {
		leClient := testAccProvider.Meta().(logentries_goclient.LogEntriesClient)
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