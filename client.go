package gx

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Client represents a client accessing a site's API
type Client struct {
	UserAgent  string
	Domain     string
	HTTPClient *http.Client
}

// New returns a Client struct configured to work with a given domain.
func New(domain string) Client {
	return Client{UserAgent: "gx", Domain: domain, HTTPClient: &http.Client{}}
}

// GetCatalog gets a catalog struct from a given board's name.
func (c *Client) GetCatalog(board string) (Catalog, error) {
	h := c.HTTPClient
	url := "https://" + c.Domain + "/" + board + "/catalog.json"

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Catalog{}, err
	}

	r.Header.Set("User-Agent", c.UserAgent)

	rs, err := h.Do(r)
	if err != nil {
		return Catalog{}, err
	}

	b, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return Catalog{}, err
	}

	var cl Catalog

	err = json.Unmarshal(b, &cl.Pages)
	if err != nil {
		return Catalog{}, err
	}

	return cl, nil

}

// GetThread gets a Thread struct from a given thread number and board name.
func (c *Client) GetThread(board string, thread int) (Thread, error) {
	h := c.HTTPClient
	url := "https://" + c.Domain + "/" + board + "/res/" + strconv.Itoa(thread) + ".json"

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Thread{}, err
	}

	r.Header.Add("User-Agent", c.UserAgent)

	rs, err := h.Do(r)
	if err != nil {
		return Thread{}, err
	}

	b, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return Thread{}, err
	}

	var t Thread

	err = json.Unmarshal(b, &t)
	if err != nil {
		return Thread{}, err
	}

	return t, nil

}
