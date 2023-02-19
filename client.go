package pronoundb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Platform string

const (
	PLATFORM_DISCORD   Platform = "discord"
	PLATFORM_GITHUB    Platform = "github"
	PLATFORM_MINECRAFT Platform = "minecraft"
	PLATFORM_TWITCH    Platform = "twitch"
	PLATFORM_TWITTER   Platform = "twitter"
)

const DEFAULT_LOCATION = `https://pronoundb.org`

type Client struct {
	HTTPClient *http.Client
	Location   string
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

func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		HTTPClient: http.DefaultClient,
		Location:   DEFAULT_LOCATION,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

type HTTPError struct {
	Status int
}

func (err HTTPError) Error() string {
	return fmt.Sprintf(`Status is %d, not 200!`, err.Status)
}

var ErrResponseNil = errors.New(`response is nil`)
var ErrBodyNil = errors.New(`response body is nil`)

// Lookup directly from pronoundb
//
// GET /api/v1/lookup?platform=[platform]&id=[id]
func (c Client) RawLookup(platform Platform, id string) (io.Reader, error) {
	q := url.Values{
		"platform": []string{string(platform)},
		"id":       []string{id},
	}

	resp, err := c.HTTPClient.Get(c.Location + "/api/v1/lookup?" + q.Encode())
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, ErrResponseNil
	}
	if resp.Body == nil {
		return nil, ErrResponseNil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			Status: resp.StatusCode,
		}
	}

	return resp.Body, nil
}

type lookupResp struct {
	Pronouns string `json:"pronouns"`
}

func (c Client) RawLookupParse(body io.Reader) (string, error) {
	resp := &lookupResp{}
	err := json.NewDecoder(body).Decode(resp)
	return resp.Pronouns, err
}

// Lookup pronoundb accounts in bulk.
//
// GET /api/v1/lookup-bulk?platform=[platform]&ids=[ids...]
func (c Client) RawLookupBulk(platform Platform, ids []string) (io.Reader, error) {
	q := url.Values{
		"platform": []string{string(platform)},
		"ids":      []string{strings.Join(ids, ",")},
	}

	resp, err := c.HTTPClient.Get(c.Location + "/api/v1/lookup-bulk?" + q.Encode())

	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, ErrResponseNil
	}
	if resp.Body == nil {
		return nil, ErrResponseNil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			Status: resp.StatusCode,
		}
	}

	return resp.Body, nil
}

func (c Client) RawLookupBulkParse(body io.Reader) (map[string]string, error) {
	m := map[string]string{}

	err := json.NewDecoder(body).Decode(&m)

	return m, err
}

func (c Client) Lookup(platform Platform, id string) (Pronoun, error) {
	resp, err := c.RawLookup(platform, id)
	if err != nil {
		return PR_THEY_THEM, err
	}
	raw, err := c.RawLookupParse(resp)
	if err != nil {
		return PR_THEY_THEM, err
	}
	pr := Pronoun(raw)
	pr.Default()
	return pr, nil
}

func (c Client) LookupBulk(platform Platform, ids []string) (map[string]Pronoun, error) {
	resp, err := c.RawLookupBulk(platform, ids)
	m := map[string]Pronoun{}

	for _, id := range ids {
		m[id] = PR_THEY_THEM
	}
	
	if err != nil {
		return m, err
	}

	raw, err := c.RawLookupBulkParse(resp)
	if err != nil {
		return m, err
	}

	for id, prRaw := range raw {
		pr := Pronoun(prRaw)
		pr.Default()
		m[id] = pr
	}

	return m, nil
}
