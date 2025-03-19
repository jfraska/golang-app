package template

type CreateTemplateRequestPayload struct {
	Title     string                   `json:"title" binding:"required,max=255"`
	Slug      string                   `json:"slug" binding:"required,max=255"`
	Thumbnail string                   `json:"thumbnail" binding:"max=255"`
	Price     int                      `json:"price" binding:"omitempty"`
	Category  string                   `json:"category" binding:"max=255"`
	Discount  int                      `json:"discount" binding:"omitempty"`
	Content   []map[string]interface{} `json:"content" binding:"required"`
	Color     []map[string]interface{} `json:"color" binding:"required"`
	Music     string                   `json:"music" binding:"required"`
}

type GetTemplateRequestPayload struct {
	Slug string `uri:"slug" binding:"required"`
}
