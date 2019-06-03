package insight

import (
	"github.com/dikhan/insight_goclient"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

var actionPeriods = []string{
	"Minute",
	"5Minute",
	"10Minute",
	"15Minute",
	"30Minute",
	"1Hour",
	"1Day",
}

func resourceInsightAction() *schema.Resource {
	return &schema.Resource{
		Create: resourceInsightActionCreate,
		Read:   resourceInsightActionRead,
		Delete: resourceInsightActionDelete,
		Update: resourceInsightActionUpdate,
		Importer: &schema.ResourceImporter{
			State: resourceInsightActionImport,
		},
		Schema: map[string]*schema.Schema{
			"min_matches_count": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  1,
			},
			"min_report_count": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  1,
			},
			"min_matches_period": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Minute",
				ValidateFunc: validation.StringInSlice(actionPeriods, false),
			},
			"min_report_period": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5Minute",
				ValidateFunc: validation.StringInSlice(actionPeriods, false),
			},
			"target_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceInsightActionCreate(data *schema.ResourceData, i interface{}) error {
	client := i.(insight_goclient.InsightClient)
	action := getInsightActionFromData(data)
	if err := client.PostAction(action); err != nil {
		return err
	}
	if err := setInsightActionData(data, action); err != nil {
		return err
	}
	return resourceInsightActionRead(data, i)
}

func resourceInsightActionImport(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightActionRead(data *schema.ResourceData, i interface{}) error {
	client := i.(insight_goclient.InsightClient)
	action, err := client.GetAction(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightActionData(data, action); err != nil {
		return err
	}
	return nil
}

func resourceInsightActionUpdate(data *schema.ResourceData, i interface{}) error {
	client := i.(insight_goclient.InsightClient)
	action := getInsightActionFromData(data)
	if err := client.PutAction(action); err != nil {
		return err
	}
	if err := setInsightActionData(data, action); err != nil {
		return err
	}
	return nil
}

func resourceInsightActionDelete(data *schema.ResourceData, i interface{}) error {
	actionId := data.Id()
	client := i.(insight_goclient.InsightClient)
	if err := client.DeleteAction(actionId); err != nil {
		return err
	}
	return nil
}

func getInsightActionFromData(data *schema.ResourceData) *insight_goclient.Action {
	var targets []insight_goclient.Target
	for _, id := range data.Get("targets").(*schema.Set).List() {
		targets = append(targets, insight_goclient.Target{Id: id.(string)})
	}
	return &insight_goclient.Action{
		Id:               data.Id(),
		MinMatchesCount:  data.Get("min_matches_count").(int),
		MinReportCount:   data.Get("min_report_count").(int),
		MinMatchesPeriod: data.Get("min_matches_period").(string),
		MinReportPeriod:  data.Get("min_report_period").(string),
		Targets:          targets,
	}
}

func setInsightActionData(data *schema.ResourceData, action *insight_goclient.Action) error {
	data.SetId(action.Id)
	data.Set("min_matches_count", action.MinMatchesCount)
	data.Set("min_report_count", action.MinReportCount)
	data.Set("min_matches_period", action.MinMatchesPeriod)
	data.Set("min_report_period", action.MinReportPeriod)
	data.Set("targets", action.Targets)
	return nil
}