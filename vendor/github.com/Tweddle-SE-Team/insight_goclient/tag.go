package insight_goclient

import (
	"encoding/json"
	"fmt"
)

const (
	TAGS_PATH = "/management/tags"
)

// The Tags resource allows you to interact with Tags in your account. The following operations are supported:
// - Get details of an existing Tag and Alert
// - Get details of a list of all Tags and Alerts
// - Create a new Tag and Alert
// - Update an existing Tag and Alert

// Tag represents the entity used to get an existing tag from the insight API
type Tag struct {
	Id       string    `json:"id,omitempty"`
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Sources  []*Source `json:"sources"`
	Actions  []*Action `json:"actions"`
	Patterns []string  `json:"patterns"`
	Labels   []*Label  `json:"labels,omitempty"`
}

// source represents the source log associated with the Tag
type Source struct {
	Id              string `json:"id"`
	Name            string `json:"name,omitempty"`
	RetentionPeriod string `json:"retention_period,omitempty"`
	StoredDays      []int  `json:"stored_days"`
}

type Tags struct {
	Tags []*Tag `json:"tags"`
}

type TagRequest struct {
	Tag *Tag `json:"tag"`
}

// GetTags gets details of an existing Tag and Alert
func (client *InsightClient) GetTags() ([]*Tag, error) {
	var tags Tags
	if err := client.get(TAGS_PATH, &tags); err != nil {
		return nil, err
	}
	return tags.Tags, nil
}

// GetTag gets details of a list of all Tags and Alerts
func (client *InsightClient) GetTag(tagId string) (*Tag, error) {
	var tagRequest TagRequest
	endpoint, err := client.getTagEndpoint(tagId)
	if err != nil {
		return nil, err
	}
	if err := client.get(endpoint, &tagRequest); err != nil {
		return nil, err
	}
	return tagRequest.Tag, nil
}

// PostTag creates a new Tag and Alert
func (client *InsightClient) PostTag(tag *Tag) error {
	tagRequest := TagRequest{tag}
	resp, err := client.post(TAGS_PATH, tagRequest)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &tagRequest)
	if err != nil {
		return err
	}
	return nil
}

// PutTag updates an existing Tag and Alert
func (client *InsightClient) PutTag(tag *Tag) error {
	tagRequest := TagRequest{tag}
	endpoint, err := client.getTagEndpoint(tag.Id)
	if err != nil {
		return err
	}
	resp, err := client.put(endpoint, tagRequest)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &tagRequest)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTag deletes a specific Tag from an account.
func (client *InsightClient) DeleteTag(tagId string) error {
	endpoint, err := client.getTagEndpoint(tagId)
	if err != nil {
		return err
	}
	return client.delete(endpoint)
}

// getTagEndPoint returns the rest end point to retrieve an individual tag
func (client *InsightClient) getTagEndpoint(tagId string) (string, error) {
	if tagId == "" {
		return "", fmt.Errorf("tagId input parameter is mandatory")
	} else {
		return fmt.Sprintf("%s/%s", TAGS_PATH, tagId), nil
	}
}
