package handlers

import (
	"github.com/SokoloSHA/ArturProject/pkg/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/data", h.getUser)
			v1.POST("/sync", h.postUser)

			user := v1.Group("/user")
			{
				user.POST("/create", h.createUser)
				user.DELETE("/:userId", h.deleteUser)
			}
		}
	}

	return router
}
