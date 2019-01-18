package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	URL "net/url"
	"strconv"
	"time"

	"github.com/rwlist/rwcore/hab"
)

type Client struct {
	Client *http.Client
	Debug  bool
}

func NewClient() Client {
	return Client{
		Client: &http.Client{Timeout: 20 * time.Second},
		Debug:  false,
	}
}

func (c Client) BaseURL() string {
	return "https://m.habr.com/kek/v1/"
}

func (c Client) Get(url string, args URL.Values) ([]byte, error) {
	if c.Debug {
		log.Printf("API Request GET :: %s :: %s", url, args)
	}
	req, err := http.NewRequest("GET", c.BaseURL()+url, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = args.Encode()
	if c.Debug {
		log.Printf("API GET URL: %s", req.URL)
	}
	return c.executeRequest(req)
}

func (c Client) FetchPageDaily(page int) (*models.PageDaily, error) {
	values := make(URL.Values)
	values.Add("date", "day")
	values.Add("sort", "date")
	values.Add("page", strconv.Itoa(page))
	b, err := c.Get("articles/", values)
	if err != nil {
		return nil, err
	}
	var result models.PageDaily
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c Client) executeRequest(req *http.Request) (response []byte, err error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		if c.Debug {
			log.Printf("API Response %s :: Error %s", resp.Status, err)
		}
		return
	}
	if c.Debug {
		log.Printf("API Response %s :: Body = %s", resp.Status, string(response))
	}
	return
}
