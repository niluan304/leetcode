package graphql

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

type Client struct {
	cookie   map[string]string
	client   *graphql.Client
	endpoint string
}

const (
	EndpointZh = "https://leetcode.cn/graphql"
	EndpointEn = "https://leetcode.com/graphql"
)

type Opt func(*Client)

func WithCookie(cookie map[string]string) Opt {
	return func(c *Client) {
		c.cookie = cookie
	}
}

func New(endpoint string, opts ...Opt) (client *Client) {
	client = &Client{
		cookie:   nil,
		client:   graphql.NewClient(endpoint),
		endpoint: endpoint,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (c *Client) request(ctx context.Context, req *graphql.Request, path string, point any) (err error) {

	cookie := ""
	for name, value := range c.cookie {
		s := fmt.Sprintf("%s=%s", name, value)
		if cookie == "" {
			cookie = s
		} else {
			cookie = fmt.Sprintf("%s; %s", cookie, s)
		}
	}

	header := map[string]string{
		"Cookie":           cookie,
		"Origin":           c.endpoint,
		"Referer":          c.endpoint + path,
		"Sec-Fetch-Dest":   "empty",
		"Sec-Fetch-Mode":   "cors",
		"Sec-Fetch-Site":   "same-site",
		"X-CSRFToken":      "undefined",
		"X-Requested-With": "XMLHttpRequest",
		"Connection":       "keep-alive",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.58",
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}

	err = c.client.Run(ctx, req, point)
	if err != nil {
		return err
	}

	return nil
}
