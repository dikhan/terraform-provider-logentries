package insight

import (
	"github.com/dikhan/insight_goclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInsightLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceInsightLogCreate,
		Read:   resourceInsightLogRead,
		Delete: resourceInsightLogDelete,
		Update: resourceInsightLogUpdate,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"logset_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"source_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"token_seed": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tokens": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"structures": {
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

func resourceInsightLogCreate(data *schema.ResourceData, i interface{}) error {
	client := i.(insight_goclient.InsightClient)
	log := getInsightLogFromData(data)
	if err := client.PostLog(log); err != nil {
		return err
	}
	if err := setInsightLogData(data, log); err != nil {
		return err
	}
	return resourceInsightLogRead(data, i)
}

func resourceInsightLogImport(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightLogRead(data *schema.ResourceData, i interface{}) error {
	client := i.(insight_goclient.InsightClient)
	log, err := client.GetLog(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightLogData(data, log); err != nil {
		return err
	}
	return nil
}

func resourceInsightLogUpdate(data *schema.ResourceData, i interface{}) error {
	client := i.(insight_goclient.InsightClient)
	log := getInsightLogFromData(data)
	if err := client.PutLog(log); err != nil {
		return err
	}
	if err := setInsightLogData(data, log); err != nil {
		return err
	}
	return nil
}

func resourceInsightLogDelete(data *schema.ResourceData, i interface{}) error {
	logId := data.Id()
	client := i.(insight_goclient.InsightClient)
	if err := client.DeleteLog(logId); err != nil {
		return err
	}
	return nil
}

func getInsightLogFromData(data *schema.ResourceData) *insight_goclient.Log {
	return &insight_goclient.Log{
		Id:          data.Id(),
		SourceType:  data.Get("source_type").(string),
		Name:        data.Get("name").(string),
		TokenSeed:   data.Get("token_seed").(string),
		Tokens:      data.Get("tokens").([]string),
		Structures:  data.Get("structures").([]string),
		LogsetsInfo: data.Get("logset_ids").([]string),
		UserData:    data.Get("user_data").(map[string]string),
	}
}

func setInsightLogData(data *schema.ResourceData, log *insight_goclient.Log) error {
	var logsets []string
	for _, logset := range log.LogsetsInfo {
		logsets := append(logsets, logset.Id)
	}
	data.SetId(log.Id)
	data.Set("name", log.Name)
	data.Set("source_type", log.SourceType)
	data.Set("token_seed", log.TokenSeed)
	data.Set("tokens", log.Tokens)
	data.Set("user_data", log.UserData)
	data.Set("structures", log.Structures)
	data.Set("logset_ids", logsets)
	return nil
}
