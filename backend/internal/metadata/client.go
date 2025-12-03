package metadata

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	EpisodesURL string
	ArcsURL     string

	Cache *Cache
	http  *http.Client
}

type Cache struct {
	EpisodesByCRC map[string]Episode
	ArcsByNumber  map[int]Arc
	LastUpdated   time.Time
}

func NewClient(episodesURL, arcsURL string) *Client {
	return &Client{
		EpisodesURL: episodesURL,
		ArcsURL:     arcsURL,
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
		Cache: &Cache{
			EpisodesByCRC: map[string]Episode{},
			ArcsByNumber:  map[int]Arc{},
		},
	}
}

func (c *Client) FetchEpisodes() (map[string]Episode, error) {
	resp, err := c.http.Get(c.EpisodesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]Episode
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) FetchArcs() ([]Arc, error) {
	resp, err := c.http.Get(c.ArcsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []Arc
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) Refresh() error {
	episodes, err := c.FetchEpisodes()
	if err != nil {
		return err
	}

	arcs, err := c.FetchArcs()
	if err != nil {
		return err
	}

	// Normalize CRC keys to uppercase
	c.Cache.EpisodesByCRC = map[string]Episode{}
	for key, ep := range episodes {
		keyUpper := strings.ToUpper(key)
		ep.File.CRC32 = strings.ToUpper(ep.File.CRC32)
		c.Cache.EpisodesByCRC[keyUpper] = ep
	}

	// Arcs
	c.Cache.ArcsByNumber = map[int]Arc{}
	for _, a := range arcs {
		c.Cache.ArcsByNumber[a.ArcNumber] = a
	}
	// log.Printf("Loaded %d episodes into metadata cache", len(c.Cache.EpisodesByCRC))
	// log.Printf("Sample key from metadata: %v", reflect.ValueOf(c.Cache.EpisodesByCRC).MapKeys())

	c.Cache.LastUpdated = time.Now()
	return nil
}
