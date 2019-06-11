package insight

import (
	"github.com/Tweddle-SE-Team/insight_goclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceInsightLogset() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLogsetRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceLogsetRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	name := data.Get("name").(string)
	logset, err := client.GetLogsetByName(name)
	if err != nil {
		return err
	}
	data.SetId(logset.Id)
	data.Set("description", logset.Description)
	return nil
}
