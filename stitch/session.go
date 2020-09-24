package stitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// PURPOSE:
// Use an API key to create a session with which to interact with the Stitch API
// Example Token:
// 		{
// 			"ephemeral_token":"<EPHEMERAL_TOKEN>"
// 		}

const HostURL string = "https://api.stitchdata.com"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

type ResponseStruct struct {
	Token string `json:"ephemeral_token"`
}

func NewSession(host, apiKey *string) (*Client, error) {
	fmt.Println("Inside NewSession()...")

	fmt.Println("Connecting to:", HostURL)
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	// Opportunity to override default hostname
	if host != nil {
		c.HostURL = *host
	}

	if apiKey != nil {
		// authenticate
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/v3/sessions/ephemeral", c.HostURL), nil)
		if err != nil {
			return nil, err
		}

		c.Token = *apiKey

		// Move this into the helper method
		//authString := fmt.Sprintf("Bearer %s", c.Token)
		//req.Header.Add("Authorization", authString)
		//req.Header.Add("Content-Type", "application/json")

		body, err := c.doRequest(req)

		fmt.Println("Response:", body)

		// parse response body
		ar := ResponseStruct{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}

		c.Token = ar.Token
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.Token)
	req.Header.Add("Content-Type", "application/json")

	// Attempt the HTTP request
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
