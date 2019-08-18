package provider

import (
	"fmt"
	"github.com/Tweddle-SE-Team/terraform-provider-insight/insight"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceInsightLabel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLabelRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"color": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceLabelRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight.InsightClient)
	name := data.Get("name").(string)
	color := data.Get("color").(string)
	labels, err := client.GetLabelsByName(name, color)
	if err != nil {
		return err
	}
	if len(labels) == 0 {
		return fmt.Errorf("No label with name %s found.", name)
	}
	if len(labels) > 1 {
		return fmt.Errorf("Multiple topics with name %s found. Please use color to get a valid label", name)
	}
	data.SetId(labels[0].Id)
	data.Set("color", labels[0].Color)
	return nil
}
