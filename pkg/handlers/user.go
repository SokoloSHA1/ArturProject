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
	id := c.Param("userId")
	err := h.service.TodoUser.DeleteUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}

func (h *Handler) getUser(c *gin.Context) {
	id := c.Query("userId")

	user, err := h.service.TodoUser.GetUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func (h *Handler) postUser(c *gin.Context) {
	var input arturproject.Params

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.service.TodoCategory.UpdateCategories(input.UpdateCategories)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.TodoItem.UpdateItems(input.UpdateItems)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.TodoTag.UpdateTags(input.UpdateTags)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.TodoItemTag.UpdateItemTags(input.UpdateItemTags)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.TodoCategory.DeleteCategories(input.User, input.DeleteCategories)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.TodoItem.DeleteItems(input.DeleteItems)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.TodoTag.DeleteTags(input.DeleteTags)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.TodoItemTag.DeleteItemTags(input.DeleteItemTags)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}

func (h *Handler) getData(c *gin.Context) {
	id := c.Query("userId")

	user, err := h.service.TodoUser.GetUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	categories, err := h.service.TodoCategory.GetCategories(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	items, err := h.service.TodoItem.GetItems(categories)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	tags, err := h.service.TodoTag.GetTags(categories)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	itemTags, err := h.service.GetItemTags(items)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		User:       user,
		Categories: categories,
		Items:      items,
		Tags:       tags,
		ItemTags:   itemTags,
	})
}
