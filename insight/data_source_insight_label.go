package insight

import (
	"fmt"
	"github.com/dikhan/insight_goclient"
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
			"label_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceLabelRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	name := data.Get("name").(string)
	color := data.Get("color").(string)
	labels, err := client.GetLabelsByName(name, color)
	if err != nil {
		return err
	}
	if len(*labels) == 0 {
		return fmt.Errorf("No label with name %s found.", name)
	}
	if len(*labels) > 1 {
		return fmt.Errorf("Multiple topics with name %s found. Please use color to get a valid label", name)
	}
	data.SetId((*labels)[0].Id)
	data.Set("color", (*labels)[0].Color)
	data.Set("label_id", (*labels)[0].Id)
	return nil
}
