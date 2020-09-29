package stitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// PURPOSE:
// Use an API key to create a session with which to interact with the Stitch API.
//
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
	// Generates an ephemeral token to create a session in the Stitch web application.
	// Ephemeral tokens expire after one hour.
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
		url := fmt.Sprintf("%s/v3/sessions/ephemeral", c.HostURL)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return nil, errors.New("Unable to connect to:" + url + " result:" + err.Error())
		}

		c.Token = *apiKey

		body, err := c.doRequest(req)

		// parse response body
		ar := ResponseStruct{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			// Invalid credentials == Empty response -> unable to parse
			return nil, errors.New("Unable to parse:" + string(body) + " result:" + err.Error())
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
