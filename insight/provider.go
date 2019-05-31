package insight

import (
	"github.com/dikhan/insight_goclient"
	"github.com/hashicorp/terraform/helper/schema"
)

// InsightProvider returns a terraform.ResourceProvider.
func InsightProvider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "API key (Read/Write) to be able to interact with Insight REST API",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Region for Insight REST API",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"insight_label": dataSourceInsightLabel(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"insight_tag":    resourceInsightTag(),
			"insight_logset": resourceInsightLogset(),
			"insight_log":    resourceInsightLog(),
			"insight_label":  resourceInsightLabel(),
			"insight_action": resourceInsightAction(),
			"insight_target": resourceInsightTarget(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("api_key").(string)
	region := d.Get("region").(string)
	client, err := insight_goclient.NewInsightClient(apiKey, region)
	if err != nil {
		return nil, err
	}
	return client, nil
}
