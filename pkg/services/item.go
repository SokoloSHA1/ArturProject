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

func (s *ItemService) DeleteItems(items []string) error {
	if len(items) == 0 {
		return nil
	}

	return s.repo.DeleteItems(items)
}

func (s *ItemService) UpdateItems(items []arturproject.Item) error {
	if len(items) == 0 {
		return nil
	}

	for _, item := range items {
		ok, err := s.repo.CheckItems(item)
		if err != nil {
			return err
		}

		if !ok {
			err = s.repo.CreateItem(item)
		} else {
			err = s.repo.UpdateItem(item)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ItemService) GetItems(categories []arturproject.Category) ([]arturproject.Item, error) {
	return s.repo.GetItems(categories)
}
