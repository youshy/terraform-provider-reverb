package reverb

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceCategories() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCategoriesRead,

		Schema: map[string]*schema.Schema{
			"categories": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString,
						},
						"full_name": {
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func dataSourceCategoriesRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
