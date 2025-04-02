package invitation

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invitation struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Subdomain string             `bson:"subdomain" json:"subdomain"`
	Published bool               `bson:"published" json:"published"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`

	TemplateID primitive.ObjectID `bson:"template_id" json:"template_id,omitempty"`
}

func NewInvitationFromCreateInvitationRequest(req CreateInvitationRequestPayload) Invitation {
	return Invitation{
		Name:      req.Name,
		Subdomain: req.Subdomain,
		Published: req.Published,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a Invitation) Validate() (err error) {
	return
}
