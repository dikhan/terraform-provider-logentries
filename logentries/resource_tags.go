package logentries

import (
	"fmt"
	"github.com/dikhan/logentries_goclient"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func tagsResource() *schema.Resource {
	return &schema.Resource{
		Create: createTag,
		Read:   readTag,
		Delete: deleteTag,
		Update: updateTag,
		Importer: &schema.ResourceImporter{
			State: importTagHelper,
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
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"actions": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_matches_count": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_report_count": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_matches_period": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"min_report_period": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"targets": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"params_set": {
										Type:     schema.TypeMap,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
									},
									"alert_content_set": {
										Type:     schema.TypeMap,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
									},
								},
							},
						},
					},
				},
				Optional: true,
			},
			"patterns": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"labels": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func createTag(data *schema.ResourceData, i interface{}) error {
	var p logentries_goclient.PostTag
	var err error

	if p, err = makePostTag(data, i); err != nil {
		return err
	}

	leClient := i.(logentries_goclient.LogEntriesClient)
	tag, err := leClient.Tags.PostTag(p)

	if err != nil {
		return err
	}
	data.SetId(tag.Id)
	return nil
}

func importTagHelper(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	name := data.Id()
	leClient := i.(logentries_goclient.LogEntriesClient)
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

func readTag(data *schema.ResourceData, i interface{}) error {
	leClient := i.(logentries_goclient.LogEntriesClient)
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

func updateTag(data *schema.ResourceData, i interface{}) error {
	var p logentries_goclient.PostTag
	var err error

	if p, err = makePostTag(data, i); err != nil {
		return err
	}

	leClient := i.(logentries_goclient.LogEntriesClient)
	tag, err := leClient.Tags.PutTag(data.Id(), p)

	if err != nil {
		return err
	}
	data.SetId(tag.Id)
	return nil
}

func deleteTag(data *schema.ResourceData, i interface{}) error {
	tagId := data.Id()
	leClient := i.(logentries_goclient.LogEntriesClient)
	if err := leClient.Tags.DeleteTag(tagId); err != nil {
		return err
	}
	return nil
}

func decodeActions(data *schema.ResourceData) ([]logentries_goclient.PostAction, error) {
	actions := []logentries_goclient.PostAction{}
	if attr, ok := data.Get("actions").([]interface{}); ok {
		for _, v := range attr {
			action := &logentries_goclient.PostAction{}
			if err := mapstructure.Decode(v, action); err != nil {
				return nil, err
			}
			actions = append(actions, *action)
		}
	}
	return actions, nil
}

func decodeSources(data *schema.ResourceData) ([]logentries_goclient.PostSource, error) {
	sources := []string{}
	if err := mapstructure.Decode(data.Get("sources").([]interface{}), &sources); err != nil {
		return nil, err
	}
	decodedSources := []logentries_goclient.PostSource{}
	for _, source := range sources {
		decodedSources = append(decodedSources, logentries_goclient.PostSource{source})
	}
	return decodedSources, nil
}

func makePostTag(data *schema.ResourceData, i interface{}) (logentries_goclient.PostTag, error) {
	var sources []logentries_goclient.PostSource
	var actions []logentries_goclient.PostAction
	var err error

	patterns := []string{}
	if err := mapstructure.Decode(data.Get("patterns").([]interface{}), &patterns); err != nil {
		return logentries_goclient.PostTag{}, err
	}

	if sources, err = decodeSources(data); err != nil {
		return logentries_goclient.PostTag{}, err
	}
	if actions, err = decodeActions(data); err != nil {
		return logentries_goclient.PostTag{}, err
	}

	leClient := i.(logentries_goclient.LogEntriesClient)
	labels := logentries_goclient.GetLabels{}

	for _, labelIdObject := range data.Get("labels").([]interface{}) {
		label := logentries_goclient.Label{}
		if label, err = leClient.Labels.GetLabel(labelIdObject.(string)); err != nil {
			return logentries_goclient.PostTag{}, err
		}
		labels = append(labels, label)
	}

	p := logentries_goclient.PostTag{
		Name:     data.Get("name").(string),
		Type:     data.Get("type").(string),
		Sources:  sources,
		Actions:  actions,
		Patterns: patterns,
		Labels:   labels,
	}
	return p, nil
}
