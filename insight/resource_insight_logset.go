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
			"log_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"user_data": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func resourceInsightLogsetCreate(data *schema.ResourceData, i interface{}) error {
	client := i.(insight_goclient.InsightClient)
	logset := getInsightLogsetFromData(data)
	if err := client.PostLogset(logset); err != nil {
		return err
	}
	if err := setInsightLogsetData(data, logset); err != nil {
		return err
	}
	return resourceInsightLogsetRead(data, i)
}

func resourceInsightLogsetImport(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightLogsetRead(data *schema.ResourceData, i interface{}) error {
	client := i.(insight_goclient.InsightClient)
	logset, err := client.GetLogset(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightLogsetData(data, logset); err != nil {
		return err
	}
	return nil
}

func resourceInsightLogsetUpdate(data *schema.ResourceData, i interface{}) error {
	client := i.(insight_goclient.InsightClient)
	logset := getInsightLogsetFromData(data)
	if err := client.PutLogset(logset); err != nil {
		return err
	}
	if err := setInsightLogsetData(data, logset); err != nil {
		return err
	}
	return nil
}

func resourceInsightLogsetDelete(data *schema.ResourceData, i interface{}) error {
	logsetId := data.Id()
	client := i.(insight_goclient.InsightClient)
	if err := client.DeleteLogset(logsetId); err != nil {
		return err
	}
	return nil
}

func getInsightLogsetFromData(data *schema.ResourceData) *insight_goclient.Logset {
	var logs []insight_goclient.Info
	for _, id := range data.Get("logs_ids").(*schema.Set).List() {
		logs = append(logs, insight_goclient.Info{Id: id.(string)})
	}
	var userData map[string]string
	if v, ok := data.GetOk("user_data"); ok {
		for key, value := range v.(map[string]interface{}) {
			userData[key] = value.(string)
		}
	}
	return &insight_goclient.Logset{
		Id:          data.Id(),
		Name:        data.Get("name").(string),
		Description: data.Get("description").(string),
		UserData:    userData,
		LogsInfo:    logs,
	}
}

func setInsightLogsetData(data *schema.ResourceData, logset *insight_goclient.Logset) error {
	var logs []string
	for _, log := range logset.LogsInfo {
		logs = append(logs, log.Id)
	}
	data.SetId(logset.Id)
	data.Set("name", logset.Name)
	data.Set("description", logset.Description)
	data.Set("user_data", logset.UserData)
	data.Set("log_ids", logs)
	return nil
}
