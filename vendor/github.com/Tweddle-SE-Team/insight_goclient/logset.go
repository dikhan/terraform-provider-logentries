package insight_goclient

import (
	"encoding/json"
	"fmt"
)

const (
	LOGSETS_PATH = "/management/logsets"
)

// The Log Set resource allows you to interact with Logs in your account. The following operations are supported:
// - Get details of an existing Log Set
// - Get details of a list of all Log Sets
// - Create a new Log Set
// - Update an existing Log Set
// - Delete a Log Set
// Structs meant for clients

// PostLogset represents the entity used to create a new logset to the insight API
type Logset struct {
	Id          string            `json:"id,omitempty"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	LogsInfo    []*Info           `json:"logs_info,omitempty"`
	UserData    map[string]string `json:"user_data,omitempty"`
}

// LogsetInfo represent information about the logset
type Info struct {
	Id    string  `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Links []*Link `json:"links,omitempty"`
}

type Link struct {
	Rel  string `json:"rel,omitempty"`
	Href string `json:"href,omitempty"`
}

type Logsets struct {
	Logsets []*Logset `json:"logsets"`
}

type LogsetRequest struct {
	Logset *Logset `json:"logset"`
}

// GetLogset gets details of a list of all Log Sets
func (client *InsightClient) GetLogsets() ([]*Logset, error) {
	var logsets Logsets
	if err := client.get(LOGSETS_PATH, &logsets); err != nil {
		return nil, err
	}
	return logsets.Logsets, nil
}

// GetLogsets gets details of an existing Log Set
func (client *InsightClient) GetLogset(logsetId string) (*Logset, error) {
	var logsetRequest LogsetRequest
	endpoint, err := client.getLogsetEndpoint(logsetId)
	if err != nil {
		return nil, err
	}
	if err := client.get(endpoint, &logsetRequest); err != nil {
		return nil, err
	}
	return logsetRequest.Logset, nil
}

func (client *InsightClient) GetLogsetByName(name string) (*Logset, error) {
	logsets, err := client.GetLogsets()
	if err != nil {
		return nil, err
	}
	for _, logset := range logsets {
		if logset.Name == name {
			return logset, nil
		}
	}
	return nil, fmt.Errorf("No logset with name %s doesn't exist", name)
}

// PostLogset creates a new LogSet
func (client *InsightClient) PostLogset(logset *Logset) error {
	logsetRequest := LogsetRequest{logset}
	resp, err := client.post(LOGSETS_PATH, logsetRequest)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &logsetRequest)
	if err != nil {
		return err
	}
	return nil
}

// PutTag updates an existing Logset
func (client *InsightClient) PutLogset(logset *Logset) error {
	logsetRequest := LogsetRequest{logset}
	endpoint, err := client.getLogsetEndpoint(logset.Id)
	if err != nil {
		return err
	}
	resp, err := client.put(endpoint, logsetRequest)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &logsetRequest)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTag deletes a specific Logset from an account.
func (client *InsightClient) DeleteLogset(logsetId string) error {
	endpoint, err := client.getLogsetEndpoint(logsetId)
	if err != nil {
		return err
	}
	return client.delete(endpoint)
}

// getLogEndpoint returns the rest end point to retrieve an individual log
func (client *InsightClient) getLogsetEndpoint(logsetId string) (string, error) {
	if logsetId == "" {
		return "", fmt.Errorf("logsetId input parameter is mandatory")
	} else {
		return fmt.Sprintf("%s/%s", LOGSETS_PATH, logsetId), nil
	}
}
