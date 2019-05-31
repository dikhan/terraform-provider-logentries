package insight

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

var apiKey = os.Getenv("INSIGHT_API_KEY")

var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProvider = InsightProvider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"insight": testAccProvider,
	}
}

func TestInsightProvider(t *testing.T) {
	if err := InsightProvider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("INGISHT_API_KEY") == "" {
		t.Fatalf("err: INSIGHT_API_KEY env variable is mandatory to run acceptance tests")
	}
}
