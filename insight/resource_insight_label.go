package insight

import (
	"fmt"
	"github.com/dikhan/insight_goclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func labelResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceInsightLabelCreate,
		Read:   resourceInsightLabelRead,
		Delete: resourceInsightLabelDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"color": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"sn": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"reserved": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceInsightLabelCreate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	label := insight_label.Label{
		Name:  data.Get("name").(string),
		Color: data.Get("color").(string),
	}
	resp, err := client.PostLabel(label)
	if err != nil {
		return err
	}
	data.SetId(resp.Id)
}

func resourceInsightLabelRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	label, err := client.GetLabel(data.Id())
	if err != nil {
		return err
	}
	data.Set("name", label.Name)
	data.Set("color", label.Color)
	data.Set("sn", label.SN)
	data.Set("reserved", label.Reserved)
	return nil
}

func resourceInsightLabelDelete(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	if err := client.DeleteLabel(data.Id()); err != nil {
		return err
	}
	return nil
}
