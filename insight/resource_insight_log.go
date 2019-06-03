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
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agent_filename": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"agent_follow": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  true,
						},
					},
				},
			},
		},
	}
}

func resourceInsightLogCreate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	log := getInsightLogFromData(data)
	if err := client.PostLog(log); err != nil {
		return err
	}
	if err := setInsightLogData(data, log); err != nil {
		return err
	}
	return resourceInsightLogRead(data, meta)
}

func resourceInsightLogImport(data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightLogRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	log, err := client.GetLog(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightLogData(data, log); err != nil {
		return err
	}
	return nil
}

func resourceInsightLogUpdate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	log := getInsightLogFromData(data)
	if err := client.PutLog(log); err != nil {
		return err
	}
	if err := setInsightLogData(data, log); err != nil {
		return err
	}
	return nil
}

func resourceInsightLogDelete(data *schema.ResourceData, meta interface{}) error {
	logId := data.Id()
	client := meta.(*insight_goclient.InsightClient)
	if err := client.DeleteLog(logId); err != nil {
		return err
	}
	return nil
}

func getInsightLogFromData(data *schema.ResourceData) *insight_goclient.Log {
	var structures []string
	var logsetIds []*insight_goclient.Info
	var tokens []string
	if v, ok := data.GetOk("structures"); ok {
		for _, structure := range v.(*schema.Set).List() {
			structures = append(structures, structure.(string))
		}
	}
	if v, ok := data.GetOk("tokens"); ok {
		for _, token := range v.(*schema.Set).List() {
			tokens = append(tokens, token.(string))
		}
	}
	if v, ok := data.GetOk("logset_ids"); ok {
		for _, logsetId := range v.(*schema.Set).List() {
			logsetIds = append(logsetIds, &insight_goclient.Info{Id: logsetId.(string)})
		}
	}
	log := insight_goclient.Log{
		Id:          data.Id(),
		SourceType:  data.Get("source_type").(string),
		Name:        data.Get("name").(string),
		TokenSeed:   data.Get("token_seed").(string),
		Tokens:      tokens,
		Structures:  structures,
		LogsetsInfo: logsetIds,
	}
	if v, ok := data.GetOk("user_data"); ok {
		userDataSettingsData := v.(*schema.Set).List()[0]
		userDataSettings := userDataSettingsData.(map[string]interface{})
		log.UserData = &insight_goclient.LogUserData{
			AgentFileName: userDataSettings["agent_filename"].(string),
			AgentFollow:   userDataSettings["agent_follow"].(string),
		}
	}
	return &log
}

func setInsightLogData(data *schema.ResourceData, log *insight_goclient.Log) error {
	var logsets []string
	for _, logset := range log.LogsetsInfo {
		logsets = append(logsets, logset.Id)
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
