package metadata

import "fmt"

func (c *Client) GetEpisodeByCRC32(crc string) (Episode, error) {
	ep, ok := c.Cache.EpisodesByCRC[crc]
	if !ok {
		return Episode{}, fmt.Errorf("episode not found for CRC %s", crc)
	}
	return ep, nil
}

func (c *Client) GetArcTitle(arcNumber int) string {
	arc, ok := c.Cache.ArcsByNumber[arcNumber]
	if !ok {
		return fmt.Sprintf("Arc %d", arcNumber)
	}
	return arc.Title
}

func (c *Client) GetArcByNumber(arcNumber int) (Arc, error) {
	arc, ok := c.Cache.ArcsByNumber[arcNumber]
	if !ok {
		return Arc{}, fmt.Errorf("arc not found for number %d", arcNumber)
	}
	return arc, nil
}
