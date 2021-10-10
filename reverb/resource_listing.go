package reverb

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/youshy/terraform-provider-reverb/client"
)

// All valid conditions accepted by the Reverb API
var conditions = []string{"nonfunctioning", "poor", "fair", "good", "verygood", "excellent", "mint", "bstock", "brandnew"}

func resourceListing() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventCreate,
		Read:   resourceEventRead,
		Update: resourceEventUpdate,

		Schema: map[string]*schema.Schema{
			"make": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model": {
				Type:     schema.TypeString,
				Required: true,
			},
			"categories": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				MaxItems: 2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"condition": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: isValidCondition,
			},
			"photos": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uri": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"finish": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"year": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sku": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"upc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"upc_does_not_apply": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"has_inventory": {
				Type:     schema.TypeString,
				Required: true,
			},
			"inventory": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"offers_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"handmade": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"shipping_profile_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"shipping": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rates": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rate": {
										Type:     schema.TypeSet,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"amount": {
													Type:     schema.TypeString,
													Required: true,
												},
												"currency": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
									"region_code": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"local": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			// Computed values
			"listing_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceEventCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Client)

	event, err := resourceEventBuild(d, meta)
	if err != nil {
		return fmt.Errorf("failed to build event: %w", err)
	}

	id, err := client.Create(event)
	if err != nil {
		return fmt.Errorf("failed to create event: %w", err)
	}

	d.SetId(id)
	d.Set("listing_id", id) // do we need to set that?

	return nil
}

func resourceEventRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Client)

	event, err := client.Read(d.Id())
	if err != nil {
		return fmt.Errorf("failed to read event: %w", err)
	}

	// TODO: To implement all fields
	d.Set("make", event.Make)

	return nil
}

func resourceEventUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Client)

	event, err := resourceEventBuild(d, meta)
	if err != nil {
		return fmt.Errorf("failed to build event: %w", err)
	}

	id, err := client.Update(d.Id(), event)
	if err != nil {
		return fmt.Errorf("failed to update event: %w", err)
	}

	d.SetId(id)

	return nil
}

func resourceEventBuild(d *schema.ResourceData, meta interface{}) (*client.Event, error) {
	productMake := d.Get("make").(string)
	model := d.Get("model").(string)
	condition := d.Get("condition").(string)
	description := d.Get("description").(string)
	finish := d.Get("finish").(string)
	year := d.Get("year").(string)
	sku := d.Get("sku").(string)
	upc := d.Get("upc").(string)
	upcDoesNotApply := d.Get("upc_does_not_apply").(string)
	hasInventory := d.Get("has_inventory").(bool)
	inventory := d.Get("inventory").(int)
	offersEnabled := d.Get("offers_enabled").(bool)
	handmade := d.Get("handmade").(bool)
	shippingProfileId := d.Get("shipping_profile_id").(int)

	var event client.Event
	event.Make = productMake
	event.Model = model
	event.Condition.UUID = condition
	event.Description = description
	event.Finish = finish
	event.Year = year
	event.Sku = sku
	event.Upc = upc
	event.UpcDoesNotApply = upcDoesNotApply
	event.HasInventory = hasInventory
	event.Inventory = inventory
	event.OffersEnabled = offersEnabled
	event.Handmade = handmade
	event.ShippingProfileId = shippingProfileId

	// parse categories
	// TODO: how to get the categories prior? how to validate them?
	categoriesRaw := d.Get("categories").(*schema.Set)
	if categoriesRaw.Len() > 0 {
		categories := make([]client.UUIDArray, categoriesRaw.Len())

		for i, v := range categoriesRaw.List() {
			m := v.(map[string]interface{})

			categories[i] = client.UUIDArray{
				UUID: m["category"].(string),
			}
		}

		event.Categories = categories
	}

	// parse photos
	photosRaw := d.Get("photos").(*schema.Set)
	if photosRaw.Len() > 0 {
		photos := make([]string, photosRaw.Len())

		for i, v := range photosRaw.List() {
			m := v.(map[string]interface{})

			photos[i] = m["uri"].(string)
		}

		event.Photos = photos
	}

	// parse shipping
	// NOTE: This is so goddamn ugly.
	shippingRaw := d.Get("shipping").(*schema.Set)
	if shippingRaw.Len() > 0 {
		var shipping client.Shipping

		for _, v := range shippingRaw.List() {
			m := v.(map[string]interface{})

			ratesRaw := m["rates"].(*schema.Set)
			rates := make([]client.Rate, ratesRaw.Len())

			for j, w := range ratesRaw.List() {
				n := w.(map[string]interface{})

				rateRaw := m["rate"].(*schema.Set)
				var (
					rate client.Price
				)

				for _, u := range rateRaw.List() {
					o := u.(map[string]interface{})

					rate = client.Price{
						Amount:   o["amount"].(string),
						Currency: o["currency"].(string),
					}
				}

				rates[j] = client.Rate{
					Rate:       rate,
					RegionCode: n["region_code"].(string),
				}
			}

			shipping = client.Shipping{
				Rates: rates,
				Local: m["local"].(bool),
			}
		}

		event.Shipping = shipping
	}

	return &event, nil
}

func isValidCondition(val interface{}, key string) ([]string, []error) {
	v := val.(string)
	var errs []error

	if ok := contains(conditions, v); !ok {
		errs = append(errs, fmt.Errorf("%s is an invalid condition, available are %s", v, conditions))
	}

	return nil, errs
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
