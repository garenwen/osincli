package osincli

import (
	"errors"
	"net/http"
	"net/url"
)

// Caches configuration URLs
type clientconfigCache struct {
	authorizeUrl *url.URL
	tokenUrl     *url.URL
}

const(
	Wechat   = "wechat"
	Github   = "github"
	LinkedIn = "linkedin"
	Normal = "normal"
)

type Client struct {
	// caches urls
	configcache clientconfigCache

	// Client configuration
	config *ClientConfig

	// Transport is the HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper

	//oauth2 server
	serverType string
	//
}

// Creates a new client

func NewClient(config *ClientConfig,server string) (*Client, error) {
	c := &Client{
		config: config,
		serverType:server,
	}
	return c, c.initialize()
}

func (c *Client) initialize() error {

	switch c.serverType {
	case Wechat:
		if c.config.Wechat.Appid == ""{
			return errors.New("Missing configuration")
		}
	case Normal:
		if c.config.ClientId == ""{
			return errors.New("Missing configuration")
		}
	}

	if c.config.AuthorizeUrl == "" ||
		c.config.TokenUrl == "" || c.config.RedirectUrl == "" {
		return errors.New("Missing configuration")
	}

	var err error

	// cache configurations
	c.configcache.authorizeUrl, err = url.Parse(c.config.AuthorizeUrl)
	if err != nil {
		return err
	}

	c.configcache.tokenUrl, err = url.Parse(c.config.TokenUrl)
	if err != nil {
		return err
	}

	return nil
}
