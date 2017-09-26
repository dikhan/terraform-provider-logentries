package logentries

import "github.com/hashicorp/terraform/helper/schema"

func tagsResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"retention_period": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"stored_days": {
							Type:     schema.TypeList,
							Elem: schema.TypeInt,
						},
					},
				},
				Optional: true,
			},
			"actions": {
				Type:     schema.TypeList,
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
										Type:     schema.TypeString,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"direct": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"teams": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"users": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
										Optional: true,
									},
									"alert_content_set": {
										Type:     schema.TypeMap,
										Elem: schema.TypeString,
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
				Type:     schema.TypeString,
				Elem: schema.TypeString,
				Optional: true,
			},
			"labels": {
				Type:     schema.TypeList,
				Elem: schema.Resource{
					Schema: labelsSchema(),
				},
				Optional: true,
			},
		},
		Create: createTag,
		Read:readTag,
		Delete: deleteTag,
		Update: updateTag,
	}
}

func createTag(data *schema.ResourceData, i interface{}) error {
	return nil
}

func readTag(data *schema.ResourceData, i interface{}) error {
	return nil
}

func updateTag(data *schema.ResourceData, i interface{}) error {
	return nil
}

func deleteTag(data *schema.ResourceData, i interface{}) error {
	return nil
}