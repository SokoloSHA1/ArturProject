package services

type TodoUser interface {
}

type TodoCategory interface {
}

type TodoItem interface {
}

type Service struct {
	TodoUser
	TodoCategory
	TodoItem
}

// func NewService(repos *repository.Repository) *Service {
// 	return &Service{}
// }
