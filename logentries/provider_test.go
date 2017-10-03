package logentries

import (
	"testing"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
)

var apiKey = os.Getenv("LOGENTRIES_API_KEY")

var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProvider = LogentriesProvider()
	testAccProviders = map[string]terraform.ResourceProvider {
		"logentries": testAccProvider,
	}
}

func TestLogentriesProvider(t *testing.T) {
	if err := LogentriesProvider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("LOGENTRIES_API_KEY") == "" {
		t.Fatalf("err: LOGENTRIES_API_KEY env variable is mandatory to run acceptance tests")
	}
}
