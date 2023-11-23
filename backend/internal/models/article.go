package models

type Article struct {
	ID string `json:"id,omitempty" bson:"_id,omitempty"`

	Title string `json:"title" bson:"title"`
	Body  string `json:"body" bson:"body"`

	TagList []string `json:"tagList,omitempty" bson:"tagList,omitempty"`

	CreatedAt string `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
