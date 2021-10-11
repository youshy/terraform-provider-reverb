package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Default Reverb API host
const HostURL string = "https://api.reverb.com/api"

type Client struct {
	HTTPClient *http.Client
	HostURL    string
	Token      string
}

func NewClient(token string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}
}

// TODO: To implement
func (c *Client) Create(event *Event) (string, error) {
	event.Condition = swapConditionToUUID(event.Condition)

	body, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	buildUri := fmt.Sprintf("%s/listings", HostURL)

	req, err := http.NewRequest(http.MethodPost, buildUri, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	res, err := c.do(req)
	if err != nil {
		return "", err
	}

	response := struct {
		Id string `json:"id"`
	}{}

	err = json.NewDecoder(bytes.NewReader(res)).Decode(&response)
	if err != nil {
		return "", err
	}

	return response.Id, nil
}

// NOTE: To be tested if it really works
func (c *Client) Read(id string) (Event, error) {
	buildUri := fmt.Sprintf("%s/listings/%s", HostURL, id)

	req, err := http.NewRequest(http.MethodGet, buildUri, nil)
	if err != nil {
		return Event{}, nil
	}

	res, err := c.do(req)
	if err != nil {
		return Event{}, nil
	}

	var e Event

	err = json.NewDecoder(bytes.NewReader(res)).Decode(&e)
	if err != nil {
		return Event{}, err
	}

	return e, nil
}

func (c *Client) Update(id string, event *Event) (string, error) {
	event.Condition = swapConditionToUUID(event.Condition)

	body, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	buildUri := fmt.Sprintf("%s/listings/%s", HostURL, id)

	req, err := http.NewRequest(http.MethodPost, buildUri, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	res, err := c.do(req)
	if err != nil {
		return "", err
	}

	response := struct {
		Id string `json:"id"`
	}{}

	err = json.NewDecoder(bytes.NewReader(res)).Decode(&response)
	if err != nil {
		return "", err
	}

	return response.Id, nil
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Accept-Version", "3.0")
	req.Header.Set("Content-Type", "application/hal+json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

func swapConditionToUUID(condition UUIDArray) UUIDArray {
	return UUIDArray{
		UUID: Conditions[condition.UUID],
	}
}
