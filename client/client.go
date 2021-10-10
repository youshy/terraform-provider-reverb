package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
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
	body, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	buildUri := fmt.Sprintf("%s/listings", HostURL)

	req, err := http.NewRequest(http.MethodPost, buildUri, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	// TODO: does this endpoint return anything?
	_, err = c.do(req)
	if err != nil {
		return "", err
	}

	// in case it doesn't
	uid, _ := uuid.NewV4()

	return uid.String(), nil
}

// TODO: To implement
func (c *Client) Read(id string) (Event, error) {

	return Event{}, nil
}

// TODO: To implement
func (c *Client) Update(id string, event *Event) (string, error) {

	return "", nil
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
