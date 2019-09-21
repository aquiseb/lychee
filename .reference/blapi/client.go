package blapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Client communicates with the http://blapi.co REST API.
type Client struct {
	base string
	http *http.Client
}

// NewClient ...
func NewClient(c *http.Client) *Client {
	// c forwards all headers set on the initial Request
	// except Authorization and Cookie
	if c == nil {
		c = http.DefaultClient
	}
	return &Client{base: "https://blapi.co/api", http: c}
}

// TODO change REST api remote request to DB request
func (c *Client) NewRequest(ctx context.Context, url string) (*http.Request, error) {
	fmt.Println("BLAPI client url ---", url)
	if len(url) == 0 {
		return nil, errors.New("invalid empty-string url")
	}

	if url[0] == '/' { // Assume the user has given a relative path.
		url = c.base + url
	}

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return r.WithContext(ctx), nil
}

// Do the request.
func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.http.Do(r)
	fmt.Println("resp, err", resp, r)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if v != nil {
		if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
			return nil, fmt.Errorf("unable to parse JSON [%s %s]: %v", r.Method, r.URL.RequestURI(), err)
		}
	}

	return resp, nil
}
