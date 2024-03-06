package controllers

import (
	"rest-api/internal/models"
	"rest-api/internal/services"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []models.Video
	Save(ctx *gin.Context) models.Video
}

type controller struct {
	service services.VideoService
}

func New(service services.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []models.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) models.Video {
	var video models.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}