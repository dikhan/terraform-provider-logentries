package insight

import (
	"github.com/Tweddle-SE-Team/insight_goclient"
	"github.com/hashicorp/terraform/helper/schema"
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
			"source_ids": {
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
			"label_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceInsightTagCreate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	tag := getInsightTagFromData(data)
	if err := client.PostTag(tag); err != nil {
		return err
	}
	if err := setInsightTagData(data, tag); err != nil {
		return err
	}
	return resourceInsightTagRead(data, meta)
}

func resourceInsightTagImport(data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightTagRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	tag, err := client.GetTag(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightTagData(data, tag); err != nil {
		return err
	}
	return nil
}

func resourceInsightTagUpdate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
	tag := getInsightTagFromData(data)
	if err := client.PutTag(tag); err != nil {
		return err
	}
	if err := setInsightTagData(data, tag); err != nil {
		return err
	}
	return nil
}

func resourceInsightTagDelete(data *schema.ResourceData, meta interface{}) error {
	tagId := data.Id()
	client := meta.(*insight_goclient.InsightClient)
	if err := client.DeleteTag(tagId); err != nil {
		return err
	}
	return nil
}

func getInsightTagFromData(data *schema.ResourceData) *insight_goclient.Tag {
	var patterns []string
	if v, ok := data.GetOk("patterns"); ok {
		for _, pattern := range v.(*schema.Set).List() {
			patterns = append(patterns, pattern.(string))
		}
	}
	var actions []*insight_goclient.Action
	for _, id := range data.Get("action_ids").(*schema.Set).List() {
		actions = append(actions, &insight_goclient.Action{Id: id.(string)})
	}
	var sources []*insight_goclient.Source
	for _, id := range data.Get("source_ids").(*schema.Set).List() {
		sources = append(sources, &insight_goclient.Source{Id: id.(string)})
	}
	var labels []*insight_goclient.Label
	for _, id := range data.Get("label_ids").(*schema.Set).List() {
		labels = append(labels, &insight_goclient.Label{Id: id.(string)})
	}
	return &insight_goclient.Tag{
		Id:       data.Id(),
		Type:     data.Get("type").(string),
		Name:     data.Get("name").(string),
		Patterns: patterns,
		Sources:  sources,
		Actions:  actions,
		Labels:   labels,
		UserData: map[string]string{"product_type": "OPS"},
	}
}

func setInsightTagData(data *schema.ResourceData, tag *insight_goclient.Tag) error {
	var labels []string
	for _, label := range tag.Labels {
		labels = append(labels, label.Id)
	}
	var sources []string
	for _, source := range tag.Sources {
		sources = append(sources, source.Id)
	}
	var actions []string
	for _, action := range tag.Actions {
		actions = append(actions, action.Id)
	}
	data.SetId(tag.Id)
	data.Set("patterns", tag.Patterns)
	data.Set("name", tag.Name)
	data.Set("type", tag.Type)
	data.Set("patterns", tag.Patterns)
	data.Set("label_ids", labels)
	data.Set("source_ids", sources)
	data.Set("action_ids", actions)
	return nil
}
