package provider

import (
	"fmt"
	"github.com/Tweddle-SE-Team/terraform-provider-insight/insight"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInsightTarget() *schema.Resource {
	return &schema.Resource{
		Create: resourceInsightTargetCreate,
		Read:   resourceInsightTargetRead,
		Delete: resourceInsightTargetDelete,
		Update: resourceInsightTargetUpdate,
		Importer: &schema.ResourceImporter{
			State: resourceInsightTargetImport,
		},
		Schema: map[string]*schema.Schema{
			"pagerduty_service_key": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"webhook_url", "slack_webhook"},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"webhook_url": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"pagerduty_service_key", "slack_webhook"},
			},
			"slack_webhook": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"pagerduty_service_key", "slack_webhook"},
			},
			"log_context": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"log_link": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceInsightTargetCreate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight.InsightClient)
	target := getInsightTargetFromData(data)
	if err := client.PostTarget(target); err != nil {
		return err
	}
	if err := setInsightTargetData(data, target); err != nil {
		return err
	}
	return resourceInsightTargetRead(data, meta)
}

func resourceInsightTargetImport(data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightTargetRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight.InsightClient)
	target, err := client.GetTarget(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightTargetData(data, target); err != nil {
		return err
	}
	return nil
}

func resourceInsightTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight.InsightClient)
	target := getInsightTargetFromData(data)
	if err := client.PutTarget(target); err != nil {
		return err
	}
	if err := setInsightTargetData(data, target); err != nil {
		return err
	}
	return nil
}

func resourceInsightTargetDelete(data *schema.ResourceData, meta interface{}) error {
	targetId := data.Id()
	client := meta.(*insight.InsightClient)
	if err := client.DeleteTarget(targetId); err != nil {
		return err
	}
	return nil
}

func getInsightTargetFromData(data *schema.ResourceData) *insight.Target {
	target := insight.Target{
		Id:   data.Id(),
		Name: data.Get("name").(string),
		AlertContentSet: &insight.TargetAlertContentSet{
			LogLink: insight.StringBool(data.Get("log_link").(bool)),
			Context: insight.StringBool(data.Get("log_context").(bool)),
		},
	}
	if attr, ok := data.GetOk("pagerduty_service_key"); ok {
		target.Type = "Pagerduty"
		target.ParameterSet = &insight.TargetParameterSet{
			ServiceKey: attr.(string),
		}
	} else if attr, ok := data.GetOk("webhook_url"); ok {
		target.Type = "Webhook"
		target.ParameterSet = &insight.TargetParameterSet{
			Url: attr.(string),
		}
	} else if attr, ok := data.GetOk("slack_webhook"); ok {
		target.Type = "Slack"
		target.ParameterSet = &insight.TargetParameterSet{
			Url: attr.(string),
		}
	}
	return &target
}

func setInsightTargetData(data *schema.ResourceData, target *insight.Target) error {
	data.SetId(target.Id)
	data.Set("log_link", target.AlertContentSet.LogLink)
	data.Set("log_context", target.AlertContentSet.Context)
	data.Set("name", target.Name)
	switch target.Type {
	case "Webhook":
		data.Set("webhook_url", target.ParameterSet.Url)
	case "Slack":
		data.Set("slack_webhook", target.ParameterSet.Url)
	case "Pagerduty":
		data.Set("pagerduty_service_key", target.ParameterSet.ServiceKey)
	default:
		return fmt.Errorf("%s target type is not supported", target.Type)
	}
	return nil
}
