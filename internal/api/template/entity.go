package template

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Template struct {
	ID        primitive.ObjectID       `bson:"_id,omitempty" json:"_id"`
	Path      string                   `bson:"path" json:"path"`
	Pricing   Pricing                  `bson:"pricing" json:"pricing"`
	Content   []map[string]interface{} `bson:"content" json:"content"`
	Color     []map[string]interface{} `bson:"color" json:"color"`
	Music     string                   `bson:"music" json:"music"`
	Meta      Meta                     `bson:"meta" json:"meta"`
	CreatedAt time.Time                `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time                `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Pricing struct {
	Price    int    `bson:"price" json:"price"`
	Discount int    `bson:"discount" json:"discount"`
	Category string `bson:"category,omitempty" json:"category,omitempty"`
}

type Meta struct {
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
	Image       string `bson:"image,omitempty" json:"image,omitempty"`
}

func NewTemplateFromCreateTemplateRequest(req CreateTemplateRequestPayload) Template {
	return Template{
		Path: req.Path,
		Pricing: Pricing{
			Price:    req.Price,
			Category: req.Category,
			Discount: req.Discount,
		},
		Content: req.Content,
		Color:   req.Color,
		Music:   req.Music,
		Meta: Meta{
			Title:       req.Title,
			Description: req.Description,
			Image:       req.Image,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a Template) Validate() (err error) {
	return
}
