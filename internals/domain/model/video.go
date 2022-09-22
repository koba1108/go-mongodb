package model

import (
	"errors"
	"fmt"
	"github.com/koba1108/go-mongodb/internals/helper"
	"time"
)

var (
	ErrVideoPlayerIdEmpty     = errors.New("player id is empty")
	ErrVideoPlayerIdInvalid   = errors.New("player id is invalid")
	ErrVideoExternalIdEmpty   = errors.New("external id is empty")
	ErrVideoNameEmpty         = errors.New("name is empty")
	ErrVideoDescriptionEmpty  = errors.New("description is empty")
	ErrVideoUploadDateEmpty   = errors.New("upload date is empty")
	ErrVideoUploadDateInvalid = errors.New("upload date is invalid")
	ErrVideoURLEmpty          = errors.New("url is empty")
	ErrVideoURLInvalid        = errors.New("url is invalid")
)

type Video struct {
	ID          string     `json:"id"`
	PlayerID    string     `json:"playerId"`
	ExternalID  string     `json:"externalId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	UploadDate  time.Time  `json:"uploadDate"`
	URL         string     `json:"url"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

func NewVideo(playerId, externalId, name, description, uploadDateStr, url string) (*Video, error) {
	if playerId == "" {
		return nil, ErrVideoPlayerIdEmpty
	}
	if helper.IsULID(playerId) == false {
		return nil, ErrVideoPlayerIdInvalid
	}
	if externalId == "" {
		return nil, ErrVideoExternalIdEmpty
	}
	if name == "" {
		return nil, ErrVideoNameEmpty
	}
	if description == "" {
		return nil, ErrVideoDescriptionEmpty
	}
	if uploadDateStr == "" {
		return nil, ErrVideoUploadDateEmpty
	}
	uploadDate, err := helper.ParseTimeFromStr(uploadDateStr, "2006-01-02")
	if err != nil {
		return nil, ErrVideoUploadDateInvalid
	}
	if url == "" {
		return nil, ErrVideoURLEmpty
	}
	if helper.IsURL(url) == false {
		return nil, ErrVideoURLInvalid
	}
	return &Video{
		ID:          helper.NewULID().String(),
		PlayerID:    playerId,
		ExternalID:  externalId,
		Name:        name,
		Description: description,
		UploadDate:  uploadDate,
		URL:         url,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func NewSamplePlayerVideos(p *Player, num int) ([]*Video, error) {
	var videos []*Video
	for i := 0; i < num; i++ {
		playerId := p.ID
		externalId := fmt.Sprintf("externalId-%d", i)
		name := fmt.Sprintf("name-%d", i)
		description := fmt.Sprintf("description-%d", i)
		uploadDateStr := fmt.Sprintf("2020-01-01")
		url := fmt.Sprintf("https://example.com/%d", i)
		v, err := NewVideo(playerId, externalId, name, description, uploadDateStr, url)
		if err != nil {
			return nil, err
		}
		videos = append(videos, v)
	}
	return videos, nil
}
