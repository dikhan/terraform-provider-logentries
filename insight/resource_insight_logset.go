package insight

import (
	"github.com/dikhan/insight_goclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInsightLogset() *schema.Resource {
	return &schema.Resource{
		Create: resourceInsightLogsetCreate,
		Read:   resourceInsightLogsetRead,
		Delete: resourceInsightLogsetDelete,
		Update: resourceInsightLogsetUpdate,
		Importer: &schema.ResourceImporter{
			State: resourceInsightLogsetImport,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceInsightLogsetCreate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	logset := getInsightLogsetFromData(data)
	if err := client.PostLogset(logset); err != nil {
		return err
	}
	if err := setInsightLogsetData(data, logset); err != nil {
		return err
	}
	return resourceInsightLogsetRead(data, meta)
}

func resourceInsightLogsetImport(data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightLogsetRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	logset, err := client.GetLogset(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightLogsetData(data, logset); err != nil {
		return err
	}
	return nil
}

func resourceInsightLogsetUpdate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	logset := getInsightLogsetFromData(data)
	logsetInfo, err := client.GetLogset(data.Id())
	if err != nil {
		return nil
	}
	logset.LogsInfo = logsetInfo.LogsInfo
	if err := client.PutLogset(logset); err != nil {
		return err
	}
	if err := setInsightLogsetData(data, logset); err != nil {
		return err
	}
	return nil
}

func resourceInsightLogsetDelete(data *schema.ResourceData, meta interface{}) error {
	logsetId := data.Id()
	client := meta.(*insight_goclient.InsightClient)
	if err := client.DeleteLogset(logsetId); err != nil {
		return err
	}
	return nil
}

func getInsightLogsetFromData(data *schema.ResourceData) *insight_goclient.Logset {
	return &insight_goclient.Logset{
		Id:          data.Id(),
		Name:        data.Get("name").(string),
		Description: data.Get("description").(string),
	}
}

func setInsightLogsetData(data *schema.ResourceData, logset *insight_goclient.Logset) error {
	data.SetId(logset.Id)
	data.Set("name", logset.Name)
	data.Set("description", logset.Description)
	return nil
}
