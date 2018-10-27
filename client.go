package main

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/proxy"
)

const (
	torSocksPort   = "socks5://127.0.0.1:9050"
	defaultTimeout = 5
)

// Client is default http client to tor proxy
type Client struct {
	httpClient *http.Client
}

// OnionSite hold information about the site
type OnionSite struct {
	Title       string
	Description string
}

// NewClient creates http client through tor proxy
func NewClient(socksPort string, timeout int) (*Client, error) {
	if len(socksPort) <= 0 {
		socksPort = torSocksPort
	}
	if timeout <= 0 {
		timeout = defaultTimeout
	}

	tbProxyURL, err := url.Parse(socksPort)
	if err != nil {
		return nil, err
	}

	tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
	if err != nil {
		return nil, err
	}

	tbTransport := &http.Transport{Dial: tbDialer.Dial}
	return &Client{
		httpClient: &http.Client{
			Transport: tbTransport,
			Timeout:   time.Duration(timeout) * time.Second,
		},
	}, nil
}

// Request execute request to onion address and returns response
// in case of timeout, returns error
func (c *Client) Request(onion Address) (*OnionSite, error) {

	resp, err := c.httpClient.Get(onion.Addr())
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	response := &OnionSite{
		Title: doc.Find("title").Text(),
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); strings.EqualFold(name, "description") {
			response.Description = s.AttrOr("content", "")
		}
	})

	return response, nil
}
