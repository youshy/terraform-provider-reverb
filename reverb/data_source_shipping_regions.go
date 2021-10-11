package reverb

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceShippingRegions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceShippingRegionsRead,

		Schema: map[string]*schema.Schema{
			"shipping_regions": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type: schema.TypeString,
						},
						"name": {
							Type: schema.TypeString,
						},
						"region_type": {
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func dataSourceShippingRegionsRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
