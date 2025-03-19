package template

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Template struct {
	ID        primitive.ObjectID       `bson:"_id,omitempty" json:"id,omitempty"`
	Title     string                   `bson:"title" json:"title"`
	Slug      string                   `bson:"slug" json:"slug"`
	Thumbnail string                   `bson:"thumbnail,omitempty" json:"thumbnail"`
	Price     int                      `bson:"price" json:"price"`
	Discount  int                      `bson:"discount" json:"discount"`
	Category  string                   `bson:"category,omitempty" json:"category"`
	Content   []map[string]interface{} `bson:"content" json:"content"`
	Color     []map[string]interface{} `bson:"color" json:"color"`
	Music     string                   `bson:"music" json:"music"`
	CreatedAt time.Time                `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time                `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

func NewTemplateFromCreateTemplateRequest(req CreateTemplateRequestPayload) Template {
	return Template{
		Title:     req.Title,
		Slug:      req.Slug,
		Thumbnail: req.Thumbnail,
		Price:     req.Price,
		Category:  req.Category,
		Discount:  req.Discount,
		Content:   req.Content,
		Color:     req.Color,
		Music:     req.Music,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a Template) Validate() (err error) {
	return
}
