package client

import (
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

	return "", nil
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
