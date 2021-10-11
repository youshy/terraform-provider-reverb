package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (c *Client) FetchConditions() (Conditions, error) {
	uri := "https://api.reverb.com/api/listing_conditions"

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return Conditions{}, err
	}

	res, err := c.do(req)
	if err != nil {
		return Conditions{}, err
	}

	var cond Conditions

	err = json.NewDecoder(bytes.NewReader(res)).Decode(&cond)
	if err != nil {
		return Conditions{}, err
	}

	return cond, nil
}
