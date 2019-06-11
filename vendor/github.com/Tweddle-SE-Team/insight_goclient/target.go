package insight_goclient

import (
	"encoding/json"
	"fmt"
)

const (
	TARGETS_PATH = "/management/targets"
)

// The Targets resource allows you to interact with Targets in your account. The following operations are supported:
// - Get details of an existing target
// - Get details of a list of all targets

// Target represents the entity used to get an existing target from the insight API
type Target struct {
	Id              string                 `json:"id,omitempty"`
	Type            string                 `json:"type,omitempty"`
	Name            string                 `json:"name,omitempty"`
	ParameterSet    *TargetParameterSet    `json:"params_set,omitempty"`
	UserData        map[string]string      `json:"user_data"`
	AlertContentSet *TargetAlertContentSet `json:"alert_content_set,omitempty"`
}

type TargetParameterSet struct {
	Url         string `json:"url,omitempty"`
	ServiceKey  string `json:"service_key,omitempty"`
	Direct      string `json:"direct,omitempty"`
	Teams       string `json:"teams,omitempty"`
	Users       string `json:"users,omitempty"`
	Description string `json:"description"`
}

type TargetAlertContentSet struct {
	LogLink StringBool `json:"le_log_link"`
	Context StringBool `json:"le_context"`
}

type Targets struct {
	Targets []*Target `json:"targets"`
}

type TargetRequest struct {
	Target *Target `json:"target"`
}

// GetTargets gets details of a list of all Targets
func (client *InsightClient) GetTargets() ([]*Target, error) {
	var targets Targets
	if err := client.get(TARGETS_PATH, &targets); err != nil {
		return nil, err
	}
	return targets.Targets, nil
}

// GetTarget gets a specific Target from an account
func (client *InsightClient) GetTarget(targetId string) (*Target, error) {
	var targetRequest TargetRequest
	endpoint, err := client.getTargetEndpoint(targetId)
	if err != nil {
		return nil, err
	}
	if err := client.get(endpoint, &targetRequest); err != nil {
		return nil, err
	}
	return targetRequest.Target, nil
}

// GetTarget gets a specific Target from an account by name
func (client *InsightClient) GetTargetsByName(name string) ([]*Target, error) {
	var result []*Target
	targets, err := client.GetTargets()
	if err != nil {
		return nil, err
	}
	for _, target := range targets {
		if target.Name == name {
			result = append(result, target)
		}
	}
	return result, nil
}

// PostTag creates a new Target
func (client *InsightClient) PostTarget(target *Target) error {
	if target.UserData == nil {
		target.UserData = make(map[string]string)
	}
	targetRequest := TargetRequest{target}
	resp, err := client.post(TARGETS_PATH, targetRequest)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &targetRequest)
	if err != nil {
		return err
	}
	return nil
}

// PutTag updates an existing Target
func (client *InsightClient) PutTarget(target *Target) error {
	if target.UserData == nil {
		target.UserData = make(map[string]string)
	}
	targetRequest := TargetRequest{target}
	endpoint, err := client.getTargetEndpoint(target.Id)
	if err != nil {
		return err
	}
	resp, err := client.put(endpoint, targetRequest)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &targetRequest)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTag deletes a specific Target from an account.
func (client *InsightClient) DeleteTarget(targetId string) error {
	endpoint, err := client.getTargetEndpoint(targetId)
	if err != nil {
		return err
	}
	return client.delete(endpoint)
}

func (client *InsightClient) getTargetEndpoint(targetId string) (string, error) {
	if targetId == "" {
		return "", fmt.Errorf("targetId input parameter is mandatory")
	} else {
		return fmt.Sprintf("%s/%s", TARGETS_PATH, targetId), nil
	}
}
