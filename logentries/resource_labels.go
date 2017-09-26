package logentries

import "github.com/hashicorp/terraform/helper/schema"

func labelsSchema() map[string]*schema.Schema{
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"sn": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"color": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"reserved": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}