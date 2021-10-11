package reverb

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceShippingRegions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceShippingRegionsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
			},
			"id": {
				Type: schema.TypeString,
			},
		},
	}
}

func dataSourceShippingRegionsRead(d *schema.ResourceData, meta interface{}) error {

}
