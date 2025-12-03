package qbittorrent

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) AddTorrent(downloadURL string) error {
	// Ensure logged in
	if c.Cookie == "" {
		if err := c.Login(); err != nil {
			return err
		}
	}

	form := url.Values{}
	form.Set("urls", downloadURL)

	req, err := http.NewRequest(
		"POST",
		c.Host+"/api/v2/torrents/add",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", c.Cookie)

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("qbittorrent add failed: %d", resp.StatusCode)
	}

	return nil
}
