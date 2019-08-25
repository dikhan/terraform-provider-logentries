package main

import (
	"github.com/Tweddle-SE-Team/terraform-provider-insight/provider"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(
		&plugin.ServeOpts{
			ProviderFunc: provider.Provider})
}
