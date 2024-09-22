package pronoundb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type HTTPError struct {
	Status int
}

func (err HTTPError) Error() string {
	return fmt.Sprintf(`Status is %d, not 200!`, err.Status)
}

// A specific error for rate limiting
// Don't worry - you can still do errors.As(&HTTPErr{}, err)
type HTTPErrorRateLimit struct {
	RetryAfter time.Duration
}

func (err HTTPErrorRateLimit) Error() string {
	return fmt.Sprintf(`Rate limited: Retry-After: %v`, err.RetryAfter)
}

func (err HTTPErrorRateLimit) Unwrap() error {
	return &HTTPError{429}
}

func (err HTTPErrorRateLimit) Is(e error) bool {
	if e, ok := e.(*HTTPError); ok {
		return e.Status == 429
	}

	return false
}
func (err HTTPErrorRateLimit) As(e any) bool {
	_, ok := e.(*HTTPError)

	return ok
}

func httpFetch[T any](method, path string, c *Client, respBody T) error {
	req, err := http.NewRequest(method, c.Location+path, nil)
	if err != nil {
		return err
	}

	for k, v := range c.Headers {
		req.Header[k] = []string{v}
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		return json.NewDecoder(resp.Body).Decode(respBody)
	}
	if resp.StatusCode == 429 {
		retry := resp.Header.Get("Retry-After")
		secs, _ := strconv.Atoi(retry)

		return &HTTPErrorRateLimit{
			RetryAfter: time.Duration(time.Duration(secs) * time.Second),
		}
	}

	return &HTTPError{resp.StatusCode}
}
