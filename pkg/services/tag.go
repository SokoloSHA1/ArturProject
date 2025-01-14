package services

import (
	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/SokoloSHA/ArturProject/pkg/repository"
)

type TagService struct {
	repo repository.TodoTag
}

func NewTagService(repo repository.TodoTag) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) DeleteTags(tags []string) error {
	if len(tags) == 0 {
		return nil
	}

	return s.repo.DeleteTags(tags)
}

func (s *TagService) UpdateTags(tags []arturproject.Tag) error {
	if len(tags) == 0 {
		return nil
	}

	for _, tag := range tags {
		ok, err := s.repo.CheckTags(tag)
		if err != nil {
			return err
		}

		if !ok {
			err = s.repo.CreateTag(tag)
		} else {
			err = s.repo.UpdateTag(tag)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *TagService) GetTags(categories []arturproject.Category) ([]arturproject.Tag, error) {
	return s.repo.GetTags(categories)
}
