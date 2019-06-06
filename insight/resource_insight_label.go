package insight

import (
	"github.com/Tweddle-SE-Team/insight_goclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInsightLabel() *schema.Resource {
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
	label := getInsightLabelFromData(data)
	if err := client.PostLabel(label); err != nil {
		return err
	}
	if err := setInsightLabelData(data, label); err != nil {
		return err
	}
	return resourceInsightLabelRead(data, meta)
}

func resourceInsightLabelImport(data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightLabelRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	label, err := client.GetLabel(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightLabelData(data, label); err != nil {
		return err
	}
	return nil
}

func resourceInsightLabelUpdate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	label := getInsightLabelFromData(data)
	if err := client.PutLabel(label); err != nil {
		return err
	}
	if err := setInsightLabelData(data, label); err != nil {
		return err
	}
	return nil
}

func resourceInsightLabelDelete(data *schema.ResourceData, meta interface{}) error {
	labelId := data.Id()
	client := meta.(*insight_goclient.InsightClient)
	if err := client.DeleteLabel(labelId); err != nil {
		return err
	}
	return nil
}

func getInsightLabelFromData(data *schema.ResourceData) *insight_goclient.Label {
	return &insight_goclient.Label{
		Id:       data.Id(),
		Color:    data.Get("color").(string),
		SN:       data.Get("sn").(int),
		Reserved: data.Get("reserved").(bool),
	}
}

func setInsightLabelData(data *schema.ResourceData, label *insight_goclient.Label) error {
	data.SetId(label.Id)
	data.Set("color", label.Color)
	data.Set("sn", label.SN)
	data.Set("reserved", label.Reserved)
	return nil
}
