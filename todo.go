package arturproject

import "github.com/google/uuid"

type Category struct {
	Id        uuid.UUID `json:"id" db:"Id"`
	UserId    uuid.UUID `json:"userId" db:"UserId"`
	Title     string    `json:"title" db:"Title"`
	OrderMode int       `json:"orderMode" db:"OrderMode"`
	CreatedAt string    `json:"createdAt" db:"CreatedAt"`
	UpdatedAt string    `json:"updatedAt" db:"UpdatedAt"`
}

type Item struct {
	Id          uuid.UUID
	CategoryId  uuid.UUID
	Title       string
	Description string
	Rating      float32
	Rank        int
	CreatedAt   string
	UpdatedAt   string
}

type Tag struct {
	Id         uuid.UUID
	CategoryId uuid.UUID
	Title      string
	CreatedAt  string
	UpdatedAt  string
}

type ItemTag struct {
	Id        string
	ItemId    uuid.UUID
	TagId     uuid.UUID
	CreatedAt string
}

type OrderMode struct {
	Id    int
	Title string
}

type Params struct {
	User             User       `json:"user"`
	UpdateCategories []Category `json:"updatedCategories"`
	UpdateItems      []Item     `json:"updatedItem"`
	UpdateItemTags   []ItemTag
	UpdateTags       []Tag
	DeleteCategories []string `json:"deletedCategories"`
	DeleteItems      []string
	DeleteItemTags   []string
	DeleteTags       []string
}
