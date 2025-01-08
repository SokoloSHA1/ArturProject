package repository

type TodoUser interface {
}

type TodoCategory interface {
}

type TodoItem interface {
}

type Repository struct {
	TodoUser
	TodoCategory
	TodoItem
}
