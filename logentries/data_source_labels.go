package logentries

import (
	"fmt"
	"time"

	"github.com/dikhan/logentries_goclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceLabels() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLabelsRead,
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

func dataSourceLabelsRead(d *schema.ResourceData, i interface{}) error {
	leClient := i.(logentries_goclient.LogEntriesClient)
	name := d.Get("name")
	color := d.Get("color")
	labels := []logentries_goclient.Label{}
	allLabels, err := leClient.Labels.GetLabels()
	if err != nil {
		return err
	}
	for _, label := range allLabels {
		if label.Name == name {
			if color != "" && label.Color != color {
				continue
			}
			labels = append(labels, label)
		}
	}

	if len(labels) == 0 {
		return fmt.Errorf("No label with name %s found.", name)
	}
	if len(labels) > 1 {
		return fmt.Errorf("Multiple topics with name %s found. Please use color to get a valid label", name)
	}

	d.SetId(time.Now().UTC().String())
	d.Set("color", labels[0].Color)
	d.Set("label_id", labels[0].Id)
	return nil
}
