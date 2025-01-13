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

	for _, category := range categories {
		ok, err := s.repo.CheckCategories(category)
		if err != nil {
			return err
		}

		if !ok {
			err = s.repo.CreateCategory(category)
		} else {
			err = s.repo.UpdateCategories(category)
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func (s *CategoryService) DeleteCategories(user arturproject.User, categories []string) error {
	if len(categories) == 0 {
		return nil
	}

	return s.repo.DeleteCategories(user, categories)
}

func (s *CategoryService) GetCategories(userId string) ([]arturproject.Category, error) {
	return s.repo.GetCategories(userId)
}
