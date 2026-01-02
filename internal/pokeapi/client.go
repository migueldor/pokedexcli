package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/migueldor/pokedexcli/internal/pokecache"
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

type LocationResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type Client struct {
	Cache      *pokecache.Cache
	HttpClient http.Client
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

func (c Client) Location(url string) (LocationResponse, error) {
	var body []byte
	var err error
	if data, ok := c.Cache.Get(url); ok {
		body = data
	} else {
		res, err := c.HttpClient.Get(url)
		if err != nil {
			return LocationResponse{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationResponse{}, err
		}
		c.Cache.Add(url, body)
	}
	locations := LocationResponse{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return LocationResponse{}, err
	}

	return locations, nil
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Cache: pokecache.NewCache(cacheInterval),
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetLocation(locationName string) (LocationResponse, error) {
	url := "https://pokeapi.co/api/v2" + "/location-area/" + locationName

	if val, ok := c.Cache.Get(url); ok {
		locationResp := LocationResponse{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationResponse{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationResponse{}, err
	}

	locationResp := LocationResponse{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationResponse{}, err
	}

	c.Cache.Add(url, dat)

	return locationResp, nil
}
