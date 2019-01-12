package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Client type is a struct for...
type Client struct {
	APIKey string
}

const (
	baseURL   = "https://api.themoviedb.org/3"
	imageURL  = "https://image.tmdb.org/t/p"
	movieURL  = "/movie/"
	tvURL     = "/tv/"
	peopleURL = "/people/"
)

func (c *Client) get(url string, data interface{}) error {

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Code (%d): Something went wrong", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(data)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) fmtOptions(o map[string]string) string {
	options := ""

	if len(o) > 0 {
		for k, v := range o {
			options += fmt.Sprintf("&%s=%s", k, v)
		}
	}

	return options
}
