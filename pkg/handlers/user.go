package handlers

import (
	"net/http"

	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {
	var input arturproject.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	err := h.service.TodoUser.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}

func (h *Handler) deleteUser(c *gin.Context) {

}

func (h *Handler) getUser(c *gin.Context) {
	var input arturproject.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	user, err := h.service.TodoUser.getUser(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
