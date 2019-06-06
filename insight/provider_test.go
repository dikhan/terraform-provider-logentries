package insight

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

var apiKey = os.Getenv("INSIGHT_API_KEY")
var region = os.Getenv("INSIGHT_REGION")

var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"insight": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("INGISHT_API_KEY") == "" {
		t.Fatalf("err: INSIGHT_API_KEY env variable is mandatory to run acceptance tests")
	}
	if os.Getenv("INGISHT_REGION") == "" {
		t.Fatalf("err: INSIGHT_REGION env variable is mandatory to run acceptance tests")
	}
}
