package insight

import (
	"fmt"
	"github.com/dikhan/insight_goclient"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
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
				Type:         schema.TypeString,
				Optional:     true,
				ConflicsWith: []string{"webhook_url", "slack_webhook"},
			},
			"webhook_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ConflicsWith: []string{"pagerduty_service_key", "slack_webhook"},
			},
			"slack_webhook": {
				Type:         schema.TypeString,
				Optional:     true,
				ConflicsWith: []string{"pagerduty_service_key", "slack_webhook"},
			},
			"log_context": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"log_link": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceInsightTargetCreate(data *schema.ResourceData, i interface{}) error {
	client := i.(insight_goclient.InsightClient)
	target, err := getInsightTargetFromData(data)
	if err != nil {
		return err
	}
	if err = client.PostTarget(&target); err != nil {
		return err
	}
	if err = setInsightTargetData(data, target); err != nil {
		return err
	}
	return resourceInsightTargetRead(data, i)
}

func resourceInsightTargetImport(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{d}, nil
}

func resourceInsightTargetRead(data *schema.ResourceData, i interface{}) error {
	client := i.(insight_goclient.InsightClient)
	target, err := client.GetTarget(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightTargetData(data, target); err != nil {
		return err
	}
	return nil
}

func resourceInsightTargetUpdate(data *schema.ResourceData, i interface{}) error {
	client := i.(insight_goclient.InsightClient)
	target, err := getInsightTargetFromData(data)
	if err != nil {
		return err
	}
	if err = client.PutTarget(&target); err != nil {
		return err
	}
	if err = setInsightTargetData(data, target); err != nil {
		return err
	}
	return nil
}

func resourceInsightTargetDelete(data *schema.ResourceData, i interface{}) error {
	targetId := data.Id()
	client := i.(insight_goclient.InsightClient)
	if err := client.DeleteTarget(targetId); err != nil {
		return err
	}
	return nil
}

func getInsightTargetFromData(data *schema.ResourceData) *insight_goclient.Target {
	target := insight_goclient.Target{
		Id: data.Id(),
		AlertContentSet: insight_client.TargetAlertContentSet{
			LogLink: data.Get("log_link").(bool),
			Context: data.Get("log_context").(bool),
		},
	}
	if attr, ok := data.GetOk("pagerduty_service_key"); ok {
		target.Type = "Pagerduty"
		target.ParameterSet = insight_client.TargetParameterSet{
			ServiceKey: attr.(string),
		}
	} else if attr, ok := data.GetOk("webhook_url"); ok {
		target.Type = "Webhook"
		target.ParameterSet = insight_client.TargetParameterSet{
			Url: attr.(string),
		}
	} else if attr, ok := data.GetOk("slack_webhook"); ok {
		target.Type = "Slack"
		target.ParameterSet = insight_client.TargetParameterSet{
			Url: attr.(string),
		}
	}
	return &target
}

func setInsightTargetData(data *schema.ResourceData, target *insight_goclient.Target) error {
	data.SetId(target.Id)
	data.Set("log_link", target.AlertContentSet.LogLink)
	data.Set("log_context", taget.AlertContentSet.LogContext)
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
