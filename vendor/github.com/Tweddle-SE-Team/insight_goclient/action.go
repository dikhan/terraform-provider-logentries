package insight_goclient

import (
	"encoding/json"
	"fmt"
)

const (
	ACTIONS_PATH = "/management/actions"
)

// The Actions resource allows you to interact with Actions in your account. The following operations are supported:
// - Get details of an existing action
// - Get details of a list of all actions

// Action represents the entity used to get an existing action from the insight API
type Action struct {
	Id               string    `json:"id,omitempty"`
	MinMatchesCount  int       `json:"min_matches_count,omitempty"`
	MinReportCount   int       `json:"min_report_count,omitempty"`
	MinMatchesPeriod string    `json:"min_matches_period,omitempty"`
	MinReportPeriod  string    `json:"min_report_period,omitempty"`
	Targets          []*Target `json:"targets,omitempty"`
	Enabled          bool      `json:"enabled,omitempty"`
	Type             string    `json:"type,omitempty"`
}

type ActionRequest struct {
	Action *Action `json:"action"`
}

type Actions struct {
	Actions []*Action `json:"actions"`
}

// GetActions gets details of a list of all Actions
func (client *InsightClient) GetActions() ([]*Action, error) {
	var actions Actions
	if err := client.get(ACTIONS_PATH, &actions); err != nil {
		return nil, err
	}
	return actions.Actions, nil
}

// GetAction gets a specific Action from an account
func (client *InsightClient) GetAction(actionId string) (*Action, error) {
	var actionRequest ActionRequest
	endpoint, err := client.getActionEndpoint(actionId)
	if err != nil {
		return nil, err
	}
	if err := client.get(endpoint, &actionRequest); err != nil {
		return nil, err
	}
	return actionRequest.Action, nil
}

// PostTag creates a new Action
func (client *InsightClient) PostAction(action *Action) error {
	actionRequest := ActionRequest{action}
	resp, err := client.post(ACTIONS_PATH, actionRequest)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &actionRequest)
	if err != nil {
		return err
	}
	return nil
}

// PutTag updates an existing Action
func (client *InsightClient) PutAction(action *Action) error {
	actionRequest := ActionRequest{action}
	endpoint, err := client.getActionEndpoint(action.Id)
	if err != nil {
		return err
	}
	resp, err := client.put(endpoint, actionRequest)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &actionRequest)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTag deletes a specific Action from an account.
func (client *InsightClient) DeleteAction(actionId string) error {
	endpoint, err := client.getActionEndpoint(actionId)
	if err != nil {
		return err
	}
	return client.delete(endpoint)
}

func (client *InsightClient) getActionEndpoint(actionId string) (string, error) {
	if actionId == "" {
		return "", fmt.Errorf("actionId input parameter is mandatory")
	} else {
		return fmt.Sprintf("%s/%s", ACTIONS_PATH, actionId), nil
	}
}
