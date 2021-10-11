package client

// Payload Event
type Event struct {
	Make              string      `json:"make"`
	Model             string      `json:"model"`
	Categories        []UUIDArray `json:"categories"`
	Condition         UUIDArray   `json:"condition"`
	Photos            []string    `json:"photos"`
	Description       string      `json:"description"`
	Finish            string      `json:"finish"`
	Price             Price       `json:"price"`
	Title             string      `json:"title"`
	Year              string      `json:"year"`
	Sku               string      `json:"sku"`
	Upc               string      `json:"upc"`
	UpcDoesNotApply   string      `json:"upc_does_not_apply"`
	HasInventory      bool        `json:"has_inventory"`
	Inventory         int         `json:"inventory"`
	OffersEnabled     bool        `json:"offers_enabled"`
	Handmade          bool        `json:"handmade"`
	ShippingProfileId int         `json:"shipping_profile_id"`
	Shipping          Shipping    `json:"shipping"`
}

type UUIDArray struct {
	UUID string `json:"uuid"`
}

type Price struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Shipping struct {
	Rates []Rate `json:"rates"`
	Local bool   `json:"local"`
}

type Rate struct {
	Rate       Price  `json:"rate"`
	RegionCode string `json:"region_code"`
}

var Conditions = map[string]string{
	"nonfunctioning": "fbf35668-96a0-4baa-bcde-ab18d6b1b329",
	"poor":           "6a9dfcad-600b-46c8-9e08-ce6e5057921e",
	"fair":           "98777886-76d0-44c8-865e-bb40e669e934",
	"good":           "f7a3f48c-972a-44c6-b01a-0cd27488d3f6",
	"verygood":       "ae4d9114-1bd7-4ec5-a4ba-6653af5ac84d",
	"excellent":      "df268ad1-c462-4ba6-b6db-e007e23922ea",
	"mint":           "ac5b9c1e-dc78-466d-b0b3-7cf712967a48",
	"bstock":         "9225283f-60c2-4413-ad18-1f5eba7a856f",
	"brandnew":       "7c3f45de-2ae0-4c81-8400-fdb6b1d74890",
}

func ConditionsOnly() []string {
	c := make([]string, 0)

	for k := range Conditions {
		c = append(c, k)
	}

	return c
}
