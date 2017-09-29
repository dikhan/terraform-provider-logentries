package logentries

import (
	"github.com/dikhan/logentries_goclient"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func tagsResource() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: labelsSchema(),
				},
				Optional: true,
			},
		},
		Create: createTag,
		Read:   readTag,
		Delete: deleteTag,
		Update: updateTag,
	}
}

func createTag(data *schema.ResourceData, i interface{}) error {

	var sources []logentries_goclient.PostSource
	var actions []logentries_goclient.PostAction
	var err error

	patterns := []string{}
	if err := mapstructure.Decode(data.Get("patterns").([]interface{}), &patterns); err != nil {
		return err
	}

	if sources, err = decodeSources(data); err != nil {
		return err
	}
	if actions, err = decodeActions(data); err != nil {
		return err
	}

	labels := logentries_goclient.GetLabels{}
	if err := mapstructure.Decode(data.Get("labels").([]interface{}), &labels); err != nil {
		return err
	}

	p := logentries_goclient.PostTag{
		Name:     data.Get("name").(string),
		Type:     data.Get("type").(string),
		Sources:  sources,
		Actions:  actions,
		Patterns: patterns,
		Labels:   labels,
	}

	leClient := i.(logentries_goclient.LogEntriesClient)
	tag, err := leClient.Tags.PostTag(p)

	if err != nil {
		return err
	}
	data.SetId(tag.Id)
	return nil
}

func readTag(data *schema.ResourceData, i interface{}) error {
	return nil
}

func updateTag(data *schema.ResourceData, i interface{}) error {
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
