package template

type CreateTemplateRequestPayload struct {
	Title       string                   `json:"title" binding:"required,max=255"`
	Description string                   `json:"description" binding:"omitempty"`
	Image       string                   `json:"image"`
	Path        string                   `json:"path" binding:"required,max=255"`
	Price       int                      `json:"price" binding:"omitempty"`
	Category    string                   `json:"category"`
	Discount    int                      `json:"discount" binding:"omitempty"`
	Content     []map[string]interface{} `json:"content" binding:"required"`
	Color       []map[string]interface{} `json:"color" binding:"required"`
	Music       string                   `json:"music" binding:"required"`
}

type GetTemplateRequestPayload struct {
	ID string `uri:"id" binding:"required"`
}
