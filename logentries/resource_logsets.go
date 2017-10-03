package logentries

import (
	"github.com/dikhan/logentries_goclient"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func logSetsResource() *schema.Resource {
	return &schema.Resource{
		Create: createLogSet,
		Read:   readLogSet,
		Delete: deleteLogSet,
		Update: updateLogSet,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"logs_info": {
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

func createLogSet(data *schema.ResourceData, i interface{}) error {
	var p logentries_goclient.PostLogSet
	var err error

	if p, err = makeLogSet(data); err != nil {
		return err
	}

	leClient := i.(logentries_goclient.LogEntriesClient)
	logSet, err := leClient.LogSets.PostLogSet(p)

	if err != nil {
		return err
	}
	data.SetId(logSet.Id)
	return nil
}

func readLogSet(data *schema.ResourceData, i interface{}) error {
	leClient := i.(logentries_goclient.LogEntriesClient)
	logSet, _, err := leClient.LogSets.GetLogSet(data.Id())

	if err != nil {
		return nil
	}

	data.Set("name", logSet.Name)
	data.Set("description", logSet.Description)
	var logsInfo []string
	for _, logInfo := range logSet.LogsInfo {
		logsInfo = append(logsInfo, logInfo.Id)
	}
	data.Set("logs_info", logsInfo)
	data.Set("user_data", logSet.UserData)
	return nil
}

func updateLogSet(data *schema.ResourceData, i interface{}) error {
	var p logentries_goclient.PutLogSet
	var err error

	leClient := i.(logentries_goclient.LogEntriesClient)
	if p, err = makePutLogSet(data, &leClient); err != nil {
		return err
	}

	logSet, err := leClient.LogSets.PutLogSet(data.Id(), p)
	if err != nil {
		return err
	}
	data.SetId(logSet.Id)
	return nil
}

func deleteLogSet(data *schema.ResourceData, i interface{}) error {
	logSetId := data.Id()
	leClient := i.(logentries_goclient.LogEntriesClient)
	if err := leClient.LogSets.DeleteLogSet(logSetId); err != nil {
		return err
	}
	return nil
}

func decodeLogsInfo(data *schema.ResourceData, fetchRemote bool, client *logentries_goclient.LogEntriesClient) ([]logentries_goclient.PostLogInfo, []logentries_goclient.LogInfo, error) {
	logsInfo := []string{}
	if err := mapstructure.Decode(data.Get("logs_info").([]interface{}), &logsInfo); err != nil {
		return nil, nil, err
	}

	decodedLogsInfo := []logentries_goclient.LogInfo{}
	decodedPostLogsInfo := []logentries_goclient.PostLogInfo{}
	for _, logId := range logsInfo {
		if fetchRemote {
			_, logInfo, err := client.Logs.GetLog(logId)
			if err != nil {
				return nil, nil, err
			}
			decodedLogsInfo = append(decodedLogsInfo, logInfo)
		} else {
			decodedPostLogsInfo = append(decodedPostLogsInfo, logentries_goclient.PostLogInfo{logId})
		}
	}
	return decodedPostLogsInfo, decodedLogsInfo, nil
}

func makeLogSet(data *schema.ResourceData) (logentries_goclient.PostLogSet, error) {
	var logsInfo []logentries_goclient.PostLogInfo
	var decodedUserData map[string]string
	var err error

	if logsInfo, _, err = decodeLogsInfo(data, false, nil); err != nil {
		return logentries_goclient.PostLogSet{}, err
	}

	if err := mapstructure.Decode(data.Get("user_data").(map[string]interface {}), &decodedUserData); err != nil {
		return logentries_goclient.PostLogSet{}, err
	}

	p := logentries_goclient.PostLogSet{
		Name: data.Get("name").(string),
		Description: data.Get("description").(string),
		UserData: decodedUserData,
		LogsInfo: logsInfo,
	}
	return p, nil
}

func makePutLogSet(data *schema.ResourceData, client *logentries_goclient.LogEntriesClient) (logentries_goclient.PutLogSet, error) {
	var decodedLogsInfo []logentries_goclient.LogInfo
	var decodedUserData map[string]string
	var err error

	if _, decodedLogsInfo, err = decodeLogsInfo(data, true, client); err != nil {
		return logentries_goclient.PutLogSet{}, err
	}

	if err := mapstructure.Decode(data.Get("user_data").(map[string]interface {}), &decodedUserData); err != nil {
		return logentries_goclient.PutLogSet{}, err
	}

	p := logentries_goclient.PutLogSet{
		Name:        data.Get("name").(string),
		Description: data.Get("description").(string),
		UserData:    decodedUserData,
		LogsInfo:    decodedLogsInfo,
	}
	return p, nil
}