package model

import (
	"errors"
	"fmt"
	"github.com/koba1108/go-mongodb/internals/helper"
)

var (
	ErrPlayerNameEmpty            = errors.New("player name is empty")
	ErrPlayerOfficialSiteUrlEmpty = errors.New("player official site url is empty")
)

type Player struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	OfficialSiteUrl string `json:"officialSiteUrl"`
}

func NewPlayer(name, officialSiteUrl string) (*Player, error) {
	if name == "" {
		return nil, ErrPlayerNameEmpty
	}
	if officialSiteUrl == "" {
		return nil, ErrPlayerOfficialSiteUrlEmpty
	}
	return &Player{
		ID:              helper.NewULID().String(),
		Name:            name,
		OfficialSiteUrl: officialSiteUrl,
	}, nil
}

func NewSamplePlayers(num int) ([]*Player, error) {
	var players []*Player
	for i := 0; i < num; i++ {
		name := fmt.Sprintf("name-%d", i)
		officialSiteUrl := fmt.Sprintf("https://example.com/%d", i)
		p, err := NewPlayer(name, officialSiteUrl)
		if err != nil {
			return nil, err
		}
		players = append(players, p)
	}
	return players, nil
}
