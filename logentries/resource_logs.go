package logentries

import (
	"github.com/dikhan/logentries_goclient"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func logsResource() *schema.Resource {
	return &schema.Resource{
		Create: createLog,
		Read:   readLog,
		Delete: deleteLog,
		Update: updateLog,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		    "logsets_info": {
				Type:     schema.TypeList,
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
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"structures": {
				Type:     schema.TypeList,
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

func createLog(data *schema.ResourceData, i interface{}) error {
	var p logentries_goclient.PostLog
	var err error

	if p, err = makeLog(data); err != nil {
		return err
	}

	leClient := i.(logentries_goclient.LogEntriesClient)
	log, err := leClient.Logs.PostLog(p)

	if err != nil {
		return err
	}
	data.Set("tokens", log.Tokens)
	data.SetId(log.Id)
	return nil
}

func readLog(data *schema.ResourceData, i interface{}) error {
	leClient := i.(logentries_goclient.LogEntriesClient)
	log, _, err := leClient.Logs.GetLog(data.Id())

	if err != nil {
		return nil
	}
	updateStateWithRemote(data, log)
	return nil
}

func updateLog(data *schema.ResourceData, i interface{}) error {
	var p logentries_goclient.PutLog
	var err error

	leClient := i.(logentries_goclient.LogEntriesClient)
	if p, err = makePutLog(data, &leClient); err != nil {
		return err
	}

	log, err := leClient.Logs.PutLog(data.Id(), p)
	if err != nil {
		return err
	}
	data.SetId(log.Id)
	return nil
}

func deleteLog(data *schema.ResourceData, i interface{}) error {
	logId := data.Id()
	leClient := i.(logentries_goclient.LogEntriesClient)
	if err := leClient.Logs.DeleteLog(logId); err != nil {
		return err
	}
	return nil
}

func updateStateWithRemote(data *schema.ResourceData, log logentries_goclient.Log) {
	data.Set("name", log.Name)
	data.Set("source_type", log.SourceType)
	data.Set("token_seed", log.TokenSeed)
	data.Set("structures", log.Structures)
	data.Set("tokens", log.Tokens)

	var logSetsInfo []string
	for _, logSetInfo := range log.LogsetsInfo {
		logSetsInfo = append(logSetsInfo, logSetInfo.Id)
	}
	data.Set("logsets_info", logSetsInfo)

	userData := map[string]string{
		"le_agent_filename": log.UserData.LogEntriesAgentFileName,
		"le_agent_follow":log.UserData.LogEntriesAgentFollow,
	}
	data.Set("user_data", userData)
}

func decodeLogSetsInfo(data *schema.ResourceData, fetchRemote bool, client *logentries_goclient.LogEntriesClient) ([]logentries_goclient.PostLogSetInfo, []logentries_goclient.LogSetInfo, error) {
	logsInfo := []string{}
	if err := mapstructure.Decode(data.Get("logsets_info").([]interface{}), &logsInfo); err != nil {
		return nil, nil, err
	}

	decodedLogSetsInfo := []logentries_goclient.LogSetInfo{}
	decodedPostLogSetsInfo := []logentries_goclient.PostLogSetInfo{}
	for _, logId := range logsInfo {
		if fetchRemote {
			_, logSet, err := client.LogSets.GetLogSet(logId)
			if err != nil {
				return nil, nil, err
			}
			decodedLogSetsInfo = append(decodedLogSetsInfo, logSet)
		} else {
			decodedPostLogSetsInfo = append(decodedPostLogSetsInfo, logentries_goclient.PostLogSetInfo{logId})
		}
	}
	return decodedPostLogSetsInfo, decodedLogSetsInfo, nil
}

func decodeUserData(data *schema.ResourceData) (logentries_goclient.LogUserData, error) {
	var decodedUserData map[string]string
	if err := mapstructure.Decode(data.Get("user_data").(map[string]interface {}), &decodedUserData); err != nil {
		return logentries_goclient.LogUserData{}, err
	}
	logUserData := logentries_goclient.LogUserData{
		LogEntriesAgentFollow: decodedUserData["le_agent_follow"],
		LogEntriesAgentFileName: decodedUserData["le_agent_filename"],
	}
	return logUserData, nil
}

func makeLog(data *schema.ResourceData) (logentries_goclient.PostLog, error) {
	var logSetsInfo []logentries_goclient.PostLogSetInfo
	var decodedUserData logentries_goclient.LogUserData
	var err error

	if logSetsInfo, _, err = decodeLogSetsInfo(data,false, nil); err != nil {
		return logentries_goclient.PostLog{}, err
	}

	structures := []string{}
	if err := mapstructure.Decode(data.Get("structures").([]interface{}), &structures); err != nil {
		return logentries_goclient.PostLog{}, err
	}

	if decodedUserData, err = decodeUserData(data); err != nil {
		return logentries_goclient.PostLog{}, err
	}

	p := logentries_goclient.PostLog{
		Name:        data.Get("name").(string),
		SourceType:  data.Get("source_type").(string),
		TokenSeed:   data.Get("token_seed").(string),
		Structures:  structures,
		LogsetsInfo: logSetsInfo,
		UserData:    decodedUserData,
	}
	return p, nil
}

func makePutLog(data *schema.ResourceData, client *logentries_goclient.LogEntriesClient) (logentries_goclient.PutLog, error) {
	var logSetsInfo []logentries_goclient.LogSetInfo
	var decodedUserData logentries_goclient.LogUserData
	var err error

	if _, logSetsInfo, err = decodeLogSetsInfo(data,true, client); err != nil {
		return logentries_goclient.PutLog{}, err
	}

	tokens := []string{}
	if err := mapstructure.Decode(data.Get("tokens").([]interface{}), &tokens); err != nil {
		return logentries_goclient.PutLog{}, err
	}

	structures := []string{}
	if err := mapstructure.Decode(data.Get("structures").([]interface{}), &structures); err != nil {
		return logentries_goclient.PutLog{}, err
	}

	if decodedUserData, err = decodeUserData(data); err != nil {
		return logentries_goclient.PutLog{}, err
	}

	p := logentries_goclient.PutLog{
		Name:        data.Get("name").(string),
		Tokens:      tokens,
		SourceType:  data.Get("source_type").(string),
		TokenSeed:   data.Get("token_seed").(string),
		Structures:  structures,
		LogsetsInfo: logSetsInfo,
		UserData:    decodedUserData,
	}
	return p, nil
}