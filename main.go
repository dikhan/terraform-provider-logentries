package terraform_provider_logentries

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/dikhan/terraform-provider-logentries/logentries"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: logentries.Provider})
}
