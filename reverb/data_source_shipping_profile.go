package reverb

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceShippingProfiles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceShippingProfilesRead,

		Schema: map[string]*schema.Schema{
			"shipping_profiles": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString,
						},
						"id": {
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func dataSourceShippingProfilesRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
