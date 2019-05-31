package main

import (
	"github.com/dikhan/terraform-provider-insight/insight"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(
		&plugin.ServeOpts{
			ProviderFunc: func() terraform.ResourceProvider {
				return insight.InsightProvider()
			},
		})
}
