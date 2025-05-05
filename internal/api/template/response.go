package template

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TemplateListResponse struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Title     string             `json:"title"`
	Image     string             `json:"image"`
	Price     int                `json:"price"`
	Category  string             `json:"category"`
	Discount  int                `json:"discount"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type TemplateDetailResponse struct {
	ID          primitive.ObjectID       `json:"id,omitempty"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Image       string                   `json:"image"`
	Path        string                   `json:"path"`
	Price       int                      `json:"price"`
	Category    string                   `json:"category"`
	Discount    int                      `json:"discount"`
	Content     []map[string]interface{} `json:"content"`
	Color       []map[string]interface{} `json:"color"`
	Music       string                   `json:"music"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
}

func NewTemplateListResponseFromEntity(templates []Template) []TemplateListResponse {
	templateList := make([]TemplateListResponse, 0)

	for _, template := range templates {

		templateList = append(templateList, TemplateListResponse{
			ID:        template.ID,
			Title:     template.Meta.Title,
			Image:     template.Meta.Image,
			Price:     template.Pricing.Price,
			Category:  template.Pricing.Category,
			Discount:  template.Pricing.Discount,
			CreatedAt: template.CreatedAt,
			UpdatedAt: template.UpdatedAt,
		})
	}

	return templateList
}

func NewTemplateDetailResponseFromEntity(template Template) TemplateDetailResponse {
	return TemplateDetailResponse{
		ID:          template.ID,
		Title:       template.Meta.Title,
		Description: template.Meta.Description,
		Image:       template.Meta.Image,
		Path:        template.Path,
		Price:       template.Pricing.Price,
		Category:    template.Pricing.Category,
		Discount:    template.Pricing.Discount,
		Content:     template.Content,
		Color:       template.Color,
		Music:       template.Music,
		CreatedAt:   template.CreatedAt,
		UpdatedAt:   template.UpdatedAt,
	}
}
