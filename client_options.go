package yuque

import "net/http"

type ClientOption func(*Client) error

// WithBaseURL sets the baseURL for the client
func WithBaseURL(urlStr string) ClientOption {
	return func(c *Client) error {
		return c.setBaseURL(urlStr)
	}
}

// WithToken sets the token for the client
func WithToken(token string) ClientOption {
	return func(c *Client) error {
		c.token = token
		return nil
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.userAgent = userAgent
		return nil
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}
