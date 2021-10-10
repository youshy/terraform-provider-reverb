package client

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
