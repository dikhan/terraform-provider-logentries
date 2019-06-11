package insight

import (
	"github.com/Tweddle-SE-Team/insight_goclient"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// InsightProvider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("INSIGHT_API_KEY", nil),
				Description: "API key (Read/Write) to be able to interact with Insight REST API",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INSIGHT_REGION", nil),
				Description: "Region for Insight REST API",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INSIGHT_ENDPOINT", nil),
				Description: "Custom Insight REST API Endpoint",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"insight_label":  dataSourceInsightLabel(),
			"insight_logset": dataSourceInsightLogset(),
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

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	apiKey := data.Get("api_key").(string)
	region := data.Get("region").(string)
	client, err := insight_goclient.NewInsightClient(apiKey, region)
	if err != nil {
		return nil, err
	}
	client.InsightUrl = data.Get("endpoint").(string)
	return client, nil
}
