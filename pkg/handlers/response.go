package handlers

import (
	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	User       arturproject.User       `json:"user"`
	Categories []arturproject.Category `json:"categories"`
	Items      []arturproject.Item     `json:"items"`
	ItemTags   []arturproject.ItemTag  `json:"itemTags"`
	Tags       []arturproject.Tag      `json:"tags"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
