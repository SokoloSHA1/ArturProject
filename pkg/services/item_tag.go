package services

import (
	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/SokoloSHA/ArturProject/pkg/repository"
)

type ItemTagService struct {
	repo repository.TodoItemTag
}

func NewItemTagService(repo repository.TodoItemTag) *ItemTagService {
	return &ItemTagService{repo: repo}
}

func (s *ItemTagService) DeleteItemTags(itemTags []string) error {
	if len(itemTags) == 0 {
		return nil
	}

	return s.repo.DeleteItemTags(itemTags)
}

func (s *ItemTagService) UpdateItemTags(itemTags []arturproject.ItemTag) error {
	return s.repo.CreateItemTags(itemTags)
}

func (s *ItemTagService) GetItemTags(items []arturproject.Item) ([]arturproject.ItemTag, error) {
	return s.repo.GetItemTags(items)
}
