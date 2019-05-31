package insight

import (
	"fmt"
	"github.com/dikhan/insight_goclient"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func resourceInsightTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceInsightTagCreate,
		Read:   resourceInsightTagRead,
		Delete: resourceInsightTagDelete,
		Update: resourceInsightTagUpdate,
		Importer: &schema.ResourceImporter{
			State: resourceInsightTagImport,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sources": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"action_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"patterns": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"labels": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceInsightTagCreate(data *schema.ResourceData, i interface{}) error {
	var p insight_goclient.PostTag
	var err error

	if p, err = makePostTag(data, i); err != nil {
		return err
	}

	leClient := i.(insight_goclient.LogEntriesClient)
	tag, err := leClient.Tags.PostTag(p)

	if err != nil {
		return err
	}
	data.SetId(tag.Id)
	return nil
}

func resourceInsightTagImport(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	name := data.Id()
	leClient := i.(insight_goclient.LogEntriesClient)
	tags, err := leClient.Tags.GetTags()
	if err != nil {
		return nil, err
	}
	for _, tag := range tags {
		if tag.Name == name {
			data.SetId(tag.Id)
			readTag(data, i)
			return []*schema.ResourceData{data}, nil
		}
	}
	return []*schema.ResourceData{}, fmt.Errorf("No tag with name %s found.", name)
}

func resourceInsightTagRead(data *schema.ResourceData, i interface{}) error {
	leClient := i.(insight_goclient.LogEntriesClient)
	tag, err := leClient.Tags.GetTag(data.Id())

	if err != nil {
		return nil
	}

	data.Set("name", tag.Name)
	data.Set("type", tag.Type)
	var sources []string
	for _, source := range tag.Sources {
		sources = append(sources, source.Id)
	}
	data.Set("sources", sources)

	actions := []map[string]interface{}{}
	for _, a := range tag.Actions {
		action := map[string]interface{}{}
		action["id"] = a.Id
		action["type"] = a.Type
		action["min_report_period"] = a.MinReportPeriod
		action["min_report_count"] = a.MinReportCount
		action["min_matches_period"] = a.MinMatchesPeriod
		action["min_matches_count"] = a.MinMatchesCount
		action["enabled"] = a.Enabled

		targets := []map[string]interface{}{}
		for _, t := range a.Targets {
			target := map[string]interface{}{}
			target["id"] = t.Id
			target["type"] = t.Type
			target["params_set"] = t.ParamsSet
			target["alert_content_set"] = t.AlertContentSet
			targets = append(targets, target)
		}
		action["targets"] = targets
		actions = append(actions, action)
	}
	data.Set("actions", actions)
	data.Set("patterns", tag.Patterns)
	data.Set("labels", tag.Labels)

	return nil
}

func resourceInsightTagUpdate(data *schema.ResourceData, i interface{}) error {
	var p insight_goclient.PostTag
	var err error

	if p, err = makePostTag(data, i); err != nil {
		return err
	}

	leClient := i.(insight_goclient.LogEntriesClient)
	tag, err := leClient.Tags.PutTag(data.Id(), p)

	if err != nil {
		return err
	}
	data.SetId(tag.Id)
	return nil
}

func resourceInsightTagDelete(data *schema.ResourceData, i interface{}) error {
	tagId := data.Id()
	leClient := i.(insight_goclient.LogEntriesClient)
	if err := leClient.Tags.DeleteTag(tagId); err != nil {
		return err
	}
	return nil
}

func decodeActions(data *schema.ResourceData) ([]insight_goclient.PostAction, error) {
	actions := []insight_goclient.PostAction{}
	if attr, ok := data.Get("actions").([]interface{}); ok {
		for _, v := range attr {
			action := &insight_goclient.PostAction{}
			if err := mapstructure.Decode(v, action); err != nil {
				return nil, err
			}
			actions = append(actions, *action)
		}
	}
	return actions, nil
}

func decodeSources(data *schema.ResourceData) ([]insight_goclient.PostSource, error) {
	sources := []string{}
	if err := mapstructure.Decode(data.Get("sources").([]interface{}), &sources); err != nil {
		return nil, err
	}
	decodedSources := []insight_goclient.PostSource{}
	for _, source := range sources {
		decodedSources = append(decodedSources, insight_goclient.PostSource{source})
	}
	return decodedSources, nil
}

func makePostTag(data *schema.ResourceData, i interface{}) (insight_goclient.PostTag, error) {
	var sources []insight_goclient.PostSource
	var actions []insight_goclient.PostAction
	var err error

	patterns := []string{}
	if err := mapstructure.Decode(data.Get("patterns").([]interface{}), &patterns); err != nil {
		return insight_goclient.PostTag{}, err
	}

	if sources, err = decodeSources(data); err != nil {
		return insight_goclient.PostTag{}, err
	}
	if actions, err = decodeActions(data); err != nil {
		return insight_goclient.PostTag{}, err
	}

	leClient := i.(insight_goclient.LogEntriesClient)
	labels := insight_goclient.GetLabels{}

	for _, labelIdObject := range data.Get("labels").([]interface{}) {
		label := insight_goclient.Label{}
		if label, err = leClient.Labels.GetLabel(labelIdObject.(string)); err != nil {
			return insight_goclient.PostTag{}, err
		}
		labels = append(labels, label)
	}

	p := insight_goclient.PostTag{
		Name:     data.Get("name").(string),
		Type:     data.Get("type").(string),
		Sources:  sources,
		Actions:  actions,
		Patterns: patterns,
		Labels:   labels,
	}
	return p, nil
}
