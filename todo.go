package arturproject

import (
	mssql "github.com/microsoft/go-mssqldb"
)

type Category struct {
	Id        mssql.UniqueIdentifier `json:"id" db:"Id"`
	UserId    mssql.UniqueIdentifier `json:"userId" db:"UserId"`
	Title     string                 `json:"title" db:"Title"`
	OrderMode int                    `json:"orderMode" db:"OrderMode"`
	CreatedAt string                 `json:"createdAt" db:"CreatedAt"`
	UpdatedAt string                 `json:"updatedAt" db:"UpdatedAt"`
}

type Item struct {
	Id          mssql.UniqueIdentifier `json:"id" db:"Id"`
	CategoryId  mssql.UniqueIdentifier `json:"CategoryId" db:"CategoryId"`
	Title       string                 `json:"Title" db:"Title"`
	Description string                 `json:"Description" db:"Description"`
	Rating      string                 `json:"Rating" db:"Rating"`
	Rank        int                    `json:"Rank" db:"Rank"`
	CreatedAt   string                 `json:"CreatedAt" db:"CreatedAt"`
	UpdatedAt   string                 `json:"UpdatedAt" db:"UpdatedAt"`
}

type Tag struct {
	Id         mssql.UniqueIdentifier `json:"id" db:"Id"`
	CategoryId mssql.UniqueIdentifier `json:"CategoryId" db:"CategoryId"`
	Title      string                 `json:"Title" db:"Title"`
	CreatedAt  string                 `json:"CreatedAt" db:"CreatedAt"`
	UpdatedAt  string                 `json:"UpdatedAt" db:"UpdatedAt"`
}

type ItemTag struct {
	Id        mssql.UniqueIdentifier `json:"id" db:"Id"`
	ItemId    mssql.UniqueIdentifier `json:"ItemId" db:"ItemId"`
	TagId     mssql.UniqueIdentifier `json:"TagId" db:"TagId"`
	CreatedAt string                 `json:"CreatedAt" db:"CreatedAt"`
}

type OrderMode struct {
	Id    int
	Title string
}

type Params struct {
	User             User       `json:"user"`
	UpdateCategories []Category `json:"updatedCategories"`
	UpdateItems      []Item     `json:"updatedItems"`
	UpdateItemTags   []ItemTag  `json:"updatedItemTags"`
	UpdateTags       []Tag      `json:"updatedTags"`
	DeleteCategories []string   `json:"deletedCategories"`
	DeleteItems      []string   `json:"deletedItems"`
	DeleteItemTags   []string   `json:"deletedItemTags"`
	DeleteTags       []string   `json:"deletedTags"`
}
