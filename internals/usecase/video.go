package usecase

import (
	"context"
	"errors"
	"github.com/koba1108/go-mongodb/internals/domain/model"
	"github.com/koba1108/go-mongodb/internals/domain/repository"
	"github.com/koba1108/go-mongodb/internals/helper"
)

type VideoUsecase interface {
	FindByID(ctx context.Context, id string) (*model.Video, error)
	FindAll(ctx context.Context) ([]*model.Video, error)
	Create(ctx context.Context, playerId, externalId, name, description, uploadDateStr, url string) (*model.Video, error)
	Update(ctx context.Context, id string, playerId, externalId, name, description, uploadDateStr, url *string) (*model.Video, error)
	Delete(ctx context.Context, id string) error
}

func NewVideoUsecase(vr repository.VideoRepository, pr repository.PlayerRepository) VideoUsecase {
	return &videoUsecase{
		videoRepository:  vr,
		playerRepository: pr,
	}
}

type videoUsecase struct {
	videoRepository  repository.VideoRepository
	playerRepository repository.PlayerRepository
}

func (vu *videoUsecase) FindByID(ctx context.Context, id string) (*model.Video, error) {
	return vu.videoRepository.FindByID(ctx, id)
}

func (vu *videoUsecase) FindAll(ctx context.Context) ([]*model.Video, error) {
	return vu.videoRepository.FindAll(ctx)
}

func (vu *videoUsecase) Create(ctx context.Context, playerId, externalId, name, description, uploadDateStr, url string) (*model.Video, error) {
	video, err := model.NewVideo(playerId, externalId, name, description, uploadDateStr, url)
	if err != nil {
		return nil, err
	}
	return vu.videoRepository.Create(ctx, video)
}

func (vu *videoUsecase) Update(ctx context.Context, id string, playerId, externalId, name, description, uploadDateStr, url *string) (*model.Video, error) {
	video, err := vu.videoRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if playerId != nil {
		player, err := vu.playerRepository.FindByID(ctx, *playerId)
		if err != nil {
			return nil, errors.New("player not found")
		}
		video.PlayerID = player.ID
	}
	if externalId != nil {
		video.ExternalID = *externalId
	}
	if name != nil {
		video.Name = *name
	}
	if description != nil {
		video.Description = *description
	}
	if uploadDateStr != nil {
		uploadDate, err := helper.ParseTimeFromStr(*uploadDateStr, "2006-01-02")
		if err != nil {
			return nil, errors.New("invalid upload date")
		}
		video.UploadDate = uploadDate
	}
	if url != nil {
		if !helper.IsURL(*url) {
			return nil, errors.New("invalid url")
		}
		video.URL = *url
	}
	return vu.videoRepository.Update(ctx, video)
}

func (vu *videoUsecase) Delete(ctx context.Context, id string) error {
	return vu.videoRepository.Delete(ctx, id)
}
