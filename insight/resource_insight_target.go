package insight

import (
	"fmt"
	"github.com/dikhan/insight_goclient"
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_link": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceInsightTargetCreate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insight_goclient.InsightClient)
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
	client := meta.(*insight_goclient.InsightClient)
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
	client := meta.(*insight_goclient.InsightClient)
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
	client := meta.(*insight_goclient.InsightClient)
	if err := client.DeleteTarget(targetId); err != nil {
		return err
	}
	return nil
}

func getInsightTargetFromData(data *schema.ResourceData) *insight_goclient.Target {
	target := insight_goclient.Target{
		Id: data.Id(),
		AlertContentSet: &insight_goclient.TargetAlertContentSet{
			LogLink: data.Get("log_link").(string),
			Context: data.Get("log_context").(string),
		},
	}
	if attr, ok := data.GetOk("pagerduty_service_key"); ok {
		target.Type = "Pagerduty"
		target.ParameterSet = &insight_goclient.TargetParameterSet{
			ServiceKey: attr.(string),
		}
	} else if attr, ok := data.GetOk("webhook_url"); ok {
		target.Type = "Webhook"
		target.ParameterSet = &insight_goclient.TargetParameterSet{
			Url: attr.(string),
		}
	} else if attr, ok := data.GetOk("slack_webhook"); ok {
		target.Type = "Slack"
		target.ParameterSet = &insight_goclient.TargetParameterSet{
			Url: attr.(string),
		}
	}
	return &target
}

func setInsightTargetData(data *schema.ResourceData, target *insight_goclient.Target) error {
	data.SetId(target.Id)
	data.Set("log_link", target.AlertContentSet.LogLink)
	data.Set("log_context", target.AlertContentSet.Context)
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
