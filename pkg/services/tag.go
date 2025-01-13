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

func (s *TagService) DeleteTags(user arturproject.User, tags []string) error {
	if len(tags) == 0 {
		return nil
	}

	return s.repo.DeleteTags(tags)
}
