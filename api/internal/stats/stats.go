package stats

import (
	"encoding/json"
	"io"
	"net/http"
)

type Player struct {
	ID       int    `json:"PlayerID"`
	Name     string `json:"Name"`
	TeamCode string `json:"Team"`
	Position string `json:"Position"`
}

type Client interface {
	GetPlayers() ([]Player, error)
}

type client struct {
	apiKey string
}

func NewClient(apiKey string) Client {
	return &client{
		apiKey: apiKey,
	}
}

func (c *client) GetPlayers() ([]Player, error) {
	response, err := http.DefaultClient.Get("https://api.sportsdata.io/v3/nfl/stats/json/PlayerSeasonStats/2023REG?key=" + c.apiKey)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var players []Player
	json.Unmarshal(responseBody, &players)

	return players, nil
}
