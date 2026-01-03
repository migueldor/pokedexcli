package pokeapi

import (
	"encoding/json"
	"io"
)

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c Client) LocationArea(url string) (LocationAreaResponse, error) {
	var body []byte
	var err error
	if data, ok := c.Cache.Get(url); ok {
		body = data
	} else {
		res, err := c.HttpClient.Get(url)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		c.Cache.Add(url, body)
	}
	locations := LocationAreaResponse{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locations, nil
}
