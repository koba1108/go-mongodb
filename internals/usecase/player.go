package usecase

import (
	"context"
	"github.com/koba1108/go-mongodb/internals/domain/model"
	"github.com/koba1108/go-mongodb/internals/domain/repository"
)

type PlayerUsecase interface {
	FindByID(ctx context.Context, id string) (*model.Player, error)
	FindByIdWithVideos(ctx context.Context, id string, limit, offset int, sortKey string, isAsc bool) ([]*model.PlayerWithVideos, error)
	FindAll(ctx context.Context) ([]*model.Player, error)
	Create(ctx context.Context, name, officialSiteUrl string) (*model.Player, error)
	Update(ctx context.Context, id string, name, officialSiteUrl *string) (*model.Player, error)
	Delete(ctx context.Context, id string) error
}

func NewPlayerUsecase(pr repository.PlayerRepository) PlayerUsecase {
	return &playerUsecase{playerRepository: pr}
}

type playerUsecase struct {
	playerRepository repository.PlayerRepository
}

func (pu *playerUsecase) FindByID(ctx context.Context, id string) (*model.Player, error) {
	return pu.playerRepository.FindByID(ctx, id)
}

func (pu *playerUsecase) FindByIdWithVideos(ctx context.Context, id string, limit, offset int, sortKey string, isAsc bool) ([]*model.PlayerWithVideos, error) {
	return pu.playerRepository.FindByIdWithVideos(ctx, id, limit, offset, sortKey, model.OrderByFromBool(isAsc))
}

func (pu *playerUsecase) FindAll(ctx context.Context) ([]*model.Player, error) {
	return pu.playerRepository.FindAll(ctx)
}

func (pu *playerUsecase) Create(ctx context.Context, name, officialSiteUrl string) (*model.Player, error) {
	player, err := model.NewPlayer(name, officialSiteUrl)
	if err != nil {
		return nil, err
	}
	return pu.playerRepository.Create(ctx, player)
}

func (pu *playerUsecase) Update(ctx context.Context, id string, name, officialSiteUrl *string) (*model.Player, error) {
	player, err := pu.playerRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if name != nil {
		player.Name = *name
	}
	if officialSiteUrl != nil {
		player.OfficialSiteUrl = *officialSiteUrl
	}
	return pu.playerRepository.Update(ctx, player)
}

func (pu *playerUsecase) Delete(ctx context.Context, id string) error {
	return pu.playerRepository.Delete(ctx, id)
}
