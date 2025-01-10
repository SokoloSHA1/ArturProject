package arturproject

type Category struct {
	Id        string
	UserId    string
	Title     string
	OrderMode int
	CreatedAt string
	UpdatedAt string
}

type Item struct {
	Id          string
	CategoryId  string
	Title       string
	Description string
	Rating      float32
	Rank        int
	CreatedAt   string
	UpdatedAt   string
}

type Tag struct {
	Id         string
	CategoryId string
	Title      string
	CreatedAt  string
	UpdatedAt  string
}

type ItemTag struct {
	Id        string
	ItemId    string
	TagId     string
	CreatedAt string
}

type OrderMode struct {
	Id    int
	Title string
}

type Params struct {
	User             User
	UpdateCategories []Category
	UpdateItems      []Item
	UpdateItemTags   []ItemTag
	UpdateTags       []Tag
	DeleteCategories []string
	DeleteItems      []string
	DeleteItemTags   []string
	DeleteTags       []string
}
