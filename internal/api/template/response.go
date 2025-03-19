package template

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TemplateListResponse struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Title     string             `json:"title"`
	Slug      string             `json:"slug"`
	Thumbnail string             `json:"thumbnail"`
	Price     int                `json:"price"`
	Category  string             `json:"category"`
	Discount  int                `json:"discount"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type TemplateDetailResponse struct {
	ID        primitive.ObjectID       `json:"id,omitempty"`
	Title     string                   `json:"title"`
	Slug      string                   `json:"slug"`
	Thumbnail string                   `json:"thumbnail"`
	Price     int                      `json:"price"`
	Category  string                   `json:"category"`
	Discount  int                      `json:"discount"`
	Content   []map[string]interface{} `json:"content"`
	Color     []map[string]interface{} `json:"color"`
	Music     string                   `json:"music"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}

func NewTemplateListResponseFromEntity(templates []Template) []TemplateListResponse {
	var templateList []TemplateListResponse

	for _, template := range templates {

		templateList = append(templateList, TemplateListResponse{
			ID:        template.ID,
			Title:     template.Title,
			Slug:      template.Slug,
			Thumbnail: template.Thumbnail,
			Price:     template.Price,
			Category:  template.Category,
			Discount:  template.Discount,
			CreatedAt: template.CreatedAt,
			UpdatedAt: template.UpdatedAt,
		})
	}

	return templateList
}

func NewTemplateDetailResponseFromEntity(template Template) TemplateDetailResponse {
	return TemplateDetailResponse{
		ID:        template.ID,
		Title:     template.Title,
		Slug:      template.Slug,
		Thumbnail: template.Thumbnail,
		Price:     template.Price,
		Category:  template.Category,
		Discount:  template.Discount,
		Content:   template.Content,
		Color:     template.Color,
		Music:     template.Music,
		CreatedAt: template.CreatedAt,
		UpdatedAt: template.UpdatedAt,
	}
}
