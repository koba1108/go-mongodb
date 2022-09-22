package handler

import (
	"github.com/koba1108/go-mongodb/internals/domain/model"
	"github.com/koba1108/go-mongodb/internals/domain/repository"
	"github.com/labstack/echo"
	"net/http"
)

type SampleDataHandler interface {
	Create(c echo.Context) error
}

func NewSampleDataHandler(vr repository.VideoRepository, pr repository.PlayerRepository) SampleDataHandler {
	return &sampleDataHandler{
		videoRepository:  vr,
		playerRepository: pr,
	}
}

type sampleDataHandler struct {
	videoRepository  repository.VideoRepository
	playerRepository repository.PlayerRepository
}

func (sdh *sampleDataHandler) Create(c echo.Context) error {
	var req struct {
		PlayerNum int `json:"playerNum"`
		VideoNum  int `json:"videoNum"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	if req.PlayerNum == 0 {
		req.PlayerNum = 10
	}
	if req.VideoNum == 0 {
		req.VideoNum = 10
	}
	players, err := model.NewSamplePlayers(req.PlayerNum)
	if err != nil {
		panic(err)
	}
	var videos []*model.Video
	for _, p := range players {
		p, err := sdh.playerRepository.Create(c.Request().Context(), p)
		if err != nil {
			panic(err)
		}
		vs, err := model.NewSamplePlayerVideos(p, req.VideoNum)
		if err != nil {
			panic(err)
		}
		for _, v := range vs {
			v, err := sdh.videoRepository.Create(c.Request().Context(), v)
			if err != nil {
				panic(err)
			}
			videos = append(videos, v)
		}
	}
	type res struct {
		Players []*model.Player `json:"players"`
		Videos  []*model.Video  `json:"videos"`
	}
	return c.JSON(http.StatusCreated, res{
		Players: players,
		Videos:  videos,
	})
}
