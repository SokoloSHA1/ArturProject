package services

import (
	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/SokoloSHA/ArturProject/pkg/repository"
)

type ItemService struct {
	repo repository.TodoItem
}

func NewItemService(repo repository.TodoItem) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) DeleteItems(user arturproject.User, items []string) error {
	if len(items) == 0 {
		return nil
	}

	return s.repo.DeleteItems(user, items)
}
