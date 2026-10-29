package httpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

type Client struct {
	http    *http.Client
	headers map[string]string
	retries int
}

type Option func(*Client)

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.http.Timeout = d
	}
}

func WithRetry(n int) Option {
	return func(c *Client) {
		c.retries = n
	}
}

func WithHeader(key, value string) Option {
	return func(c *Client) {
		c.headers[key] = value
	}
}

func New(opts ...Option) *Client {
	c := &Client{
		http:    &http.Client{Timeout: 30 * time.Second},
		headers: map[string]string{},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) Get(ctx context.Context, url string) (*http.Response, error) {
	return c.do(ctx, http.MethodGet, url, nil)
}

func (c *Client) Post(ctx context.Context, url string, body []byte) (*http.Response, error) {
	return c.do(ctx, http.MethodPost, url, body)
}

func (c *Client) Put(ctx context.Context, url string, body []byte) (*http.Response, error) {
	return c.do(ctx, http.MethodPut, url, body)
}

func (c *Client) Delete(ctx context.Context, url string) (*http.Response, error) {
	return c.do(ctx, http.MethodDelete, url, nil)
}

func (c *Client) do(ctx context.Context, method, url string, body []byte) (*http.Response, error) {
	attempts := c.retries + 1
	var (
		resp *http.Response
		err  error
	)

	for i := 0; i < attempts; i++ {
		var r io.Reader
		if body != nil {
			r = bytes.NewReader(body)
		}

		var req *http.Request
		req, err = http.NewRequestWithContext(ctx, method, url, r)
		if err != nil {
			return nil, err
		}

		for k, v := range c.headers {
			req.Header.Set(k, v)
		}

		resp, err = c.http.Do(req)
		if err == nil {
			return resp, nil
		}

		if ctx.Err() != nil {
			return nil, err
		}
	}

	return nil, err
}
