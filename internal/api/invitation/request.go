package invitation

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateInvitationRequestPayload struct {
	Name      string `json:"name" binding:"required"`
	Subdomain string `json:"subdomain" binding:"required"`
	Published bool   `json:"published" binding:"omitempty"`

	TemplateID primitive.ObjectID `json:"template_id" binding:"omitempty"`
}
