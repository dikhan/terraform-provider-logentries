package logentries

import (
	"github.com/dikhan/logentries_goclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func labelsResource() *schema.Resource {
	return &schema.Resource{
		Create: createLabel,
		Read:   readLabel,
		Delete: deleteLabel,
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

func createLabel(data *schema.ResourceData, i interface{}) error {
	var p logentries_goclient.PostLabel
	var err error

	if p, err = makeLabel(data); err != nil {
		return err
	}

	leClient := i.(logentries_goclient.LogEntriesClient)
	label, err := leClient.Labels.PostLabel(p)

	if err != nil {
		return err
	}
	data.SetId(label.Id)
	return nil
}

func readLabel(data *schema.ResourceData, i interface{}) error {
	leClient := i.(logentries_goclient.LogEntriesClient)
	label, err := leClient.Labels.GetLabel(data.Id())

	if err != nil {
		return nil
	}
	data.Set("name", label.Name)
	data.Set("color", label.Color)
	data.Set("sn", label.SN)
	data.Set("reserved", label.Reserved)
	return nil
}

func deleteLabel(data *schema.ResourceData, i interface{}) error {
	labelId := data.Id()
	leClient := i.(logentries_goclient.LogEntriesClient)
	if err := leClient.Labels.DeleteLabel(labelId); err != nil {
		return err
	}
	return nil
}

func makeLabel(data *schema.ResourceData) (logentries_goclient.PostLabel, error) {
	p := logentries_goclient.PostLabel{
		Name:  data.Get("name").(string),
		Color: data.Get("color").(string),
	}
	return p, nil
}
