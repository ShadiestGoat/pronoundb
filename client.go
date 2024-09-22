package pronoundb

import (
	"net/http"
)

const (
	DEFAULT_LOCATION     = `https://pronoundb.org`
	BULK_LOOKUP_ID_LIMIT = 50
	LIB_USER_AGENT       = `(lib:ShadiestGoat/pronoundb)`
)

type Client struct {
	HTTPClient *http.Client
	Location   string
	Headers    map[string]string
}

type ClientOption func(c *Client)

// Option to make a client have a custom location. Location should be a url, like `https://pronoundb.org`. Do not include a trailing slash
func WithCustomLocation(location string) ClientOption {
	return func(c *Client) {
		c.Location = location
	}
}

// Option to make a client have a custom http client
func WithCustomHTTPClient(httpC *http.Client) ClientOption {
	return func(c *Client) {
		c.HTTPClient = httpC
	}
}

// Option to add a bunch of headers to each request
func WithCustomHeaders(h map[string]string) ClientOption {
	return func(c *Client) {
		for k, v := range h {
			c.Headers[k] = v
		}
	}
}

// Each key is optional, but it is recommended to have all keys
type UserAgent struct {
	App, Version, Site string
}

// Option to add a user agent to each request
// This is a wrapper around WithCustomHeaders
func WithUserAgent(ua UserAgent) ClientOption {
	h := ""

	for _, v := range []*string{&ua.App, &ua.Version, &ua.Site} {
		if *v == "" {
			continue
		}
		h += *v + " "
	}

	if len(h) == 0 {
		return nil
	} else {
		h += LIB_USER_AGENT
	}

	return WithCustomHeaders(map[string]string{"User-Agent": h})
}

// Create a new client for the pronoundb v2 api
// Opts are optional but the `WithUserAgent` is highly recommended, as it can help you avoid rate limits
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		HTTPClient: http.DefaultClient,
		Location:   DEFAULT_LOCATION,
		Headers:    map[string]string{},
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}

		opt(c)
	}

	return c
}
