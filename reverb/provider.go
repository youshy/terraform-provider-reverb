package reverb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/youshy/terraform-provider-reverb/client"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"reverb_listing": resourceListing(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"reverb_categories":        dataSourceCategories(),
			"reverb_condition":         dataSourceCondition(),
			"reverb_shipping_profiles": dataSourceShippingProfiles(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		return providerConfigure(d)
	}

	return provider
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	token := d.Get("token").(string)

	c := client.NewClient(token)

	return c, nil
}
