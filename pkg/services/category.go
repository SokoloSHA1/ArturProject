package services

import (
	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/SokoloSHA/ArturProject/pkg/repository"
)

type CategoryService struct {
	repo repository.TodoCategory
}

func NewCategoryService(repo repository.TodoCategory) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) UpdateCategories(categories []arturproject.Category) error {
	if len(categories) == 0 {
		return nil
	}

	return s.repo.UpdateCategories(categories)
}

func (s *CategoryService) DeleteCategories(cagories []string) error {
	if len(cagories) == 0 {
		return nil
	}

	return s.repo.DeleteCategories(cagories)
}
