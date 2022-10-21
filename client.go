package gx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client represents a client accessing a site's API
type Client struct {
	UserAgent  string
	Domain     string
	HTTPClient *http.Client
}

// New returns a Client struct configured to work with a given domain.
func New(domain string) *Client {
	return &Client{UserAgent: "gx", Domain: domain, HTTPClient: &http.Client{}}
}

// GetCatalog gets a catalog struct from a given board's name.
func (c *Client) GetCatalog(board string) (*Catalog, error) {
	url := fmt.Sprintf("https://%s/%s/catalog.json", c.Domain, board)

	b, err := c.getRead(url)
	if err != nil {
		return nil, err
	}

	var cl Catalog

	err = json.Unmarshal(b, &cl.Pages)
	if err != nil {
		return nil, err
	}

	return &cl, nil

}

// GetThread gets a Thread struct from a given thread number and board name.
func (c *Client) GetThread(board string, number uint64) (*Thread, error) {
	url := fmt.Sprintf("https://%s/%s/res/%d.json", c.Domain, board, number)
	b, err := c.getRead(url)
	if err != nil {
		return nil, err
	}

	var t *Thread

	err = json.Unmarshal(b, t)
	if err != nil {
		return nil, err
	}

	return t, nil

}

func (c *Client) getRead(url string) ([]byte, error) {
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	return io.ReadAll(resp)

}

func (c *Client) get(url string) (io.ReadCloser, error) {
	h := c.HTTPClient

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("User-Agent", c.UserAgent)

	response, err := h.Do(request)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}
