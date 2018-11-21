package logentries

import (
	"github.com/dikhan/logentries_goclient"
	"github.com/hashicorp/terraform/helper/schema"
)

// LogentriesProvider returns a terraform.ResourceProvider.
func LogentriesProvider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Api key (Read/Write) to be able to interact with Logentries REST API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"logentries_tags":    tagsResource(),
			"logentries_logsets": logSetsResource(),
			"logentries_logs":    logsResource(),
			"logentries_labels":  labelsResource(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("api_key").(string)
	leClient, err := logentries_goclient.NewLogEntriesClient(apiKey)
	if err != nil {
		return nil, err
	}
	return leClient, nil
}
