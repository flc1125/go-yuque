package yuque

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL   = "https://www.yuque.com/api/v2/"
	defaultUserAgent = "go-yuque"
)

var defaultHTTPClient = NewRetryableHTTPClient()

type Client struct {
	// baseURL for API requests.
	baseURL *url.URL

	// token used for authentication
	token string

	// userAgent used for HTTP requests
	userAgent string

	// httpClient is the HTTP client used to communicate with the API.
	httpClient *http.Client

	// services used for talking to different parts of the Tapd API.
	UserService      *userService
	DocService       *docService
	StatisticService *statisticService
}

// NewClient returns a new Tapd API client.
// Alias for NewBasicAuthClient.
func NewClient(token string, opts ...ClientOption) (*Client, error) {
	return newClient(append(opts, WithToken(token))...)
}

// newClient returns a new Tapd API client.
func newClient(opts ...ClientOption) (*Client, error) {
	c := &Client{
		userAgent:  defaultUserAgent,
		httpClient: defaultHTTPClient,
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	// setup
	if err := c.setup(); err != nil {
		return nil, err
	}

	// services
	c.UserService = &userService{c}
	c.DocService = &docService{c}
	c.StatisticService = &statisticService{c}

	return c, nil
}

// setup sets up the client for API requests.
func (c *Client) setup() error {
	if c.baseURL == nil {
		if err := c.setBaseURL(defaultBaseURL); err != nil {
			return err
		}
	}

	return nil
}

// setBaseURL sets the base URL for API requests to a custom endpoint.
func (c *Client) setBaseURL(urlStr string) error {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	c.baseURL = baseURL

	return nil
}

func (c *Client) NewRequest(ctx context.Context, method, path string, data any, opts []RequestOption) (*http.Request, error) { //nolint:lll
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	// Set the encoded path data
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	if c.userAgent != "" {
		reqHeaders.Set("User-Agent", c.userAgent)
	}

	var body io.Reader
	switch {
	case method == http.MethodPatch || method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		if data != nil {
			b, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			body = io.NopCloser(bytes.NewReader(b))
		}
	case data != nil:
		q, err := query.Values(data)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// Token authentication
	if c.token != "" {
		req.Header.Set("X-Auth-Token", c.token)
	}

	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	// Apply request options
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v any) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()              // nolint:errcheck
	defer io.Copy(io.Discard, resp.Body) // nolint:errcheck

	// decode response body
	var rawBody RawBody
	if err := json.NewDecoder(resp.Body).Decode(&rawBody); err != nil {
		return nil, err
	}

	// check status
	if resp.StatusCode != http.StatusOK {
		return nil, &ErrorResponse{
			response: resp,
			rawBody:  &rawBody,
			err:      errors.New(rawBody.Message),
		}
	}

	if v != nil {
		if err := json.Unmarshal(rawBody.Data, v); err != nil {
			return nil, err
		}
	}

	return &Response{Response: resp, rawBody: &rawBody}, nil
}
