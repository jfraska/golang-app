package invitation

type CreateInvitationRequestPayload struct {
	Name      string `json:"name" binding:"required"`
	Subdomain string `json:"subdomain" binding:"required"`
	Published bool   `json:"published" binding:"omitempty"`

	TemplateID string `json:"template_id" binding:"omitempty"`
	UserID     string `json:"user_id" binding:"omitempty"`
}
