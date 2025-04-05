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

	TemplateIDs []primitive.ObjectID `bson:"template_ids" json:"templateIds,omitempty"`
	UserIDs     []primitive.ObjectID `bson:"user_ids" json:"userIds,omitempty"`
}

func NewInvitationFromCreateInvitationRequest(req CreateInvitationRequestPayload) Invitation {
	newTemplateID, _ := primitive.ObjectIDFromHex(req.TemplateID)
	newUserID, _ := primitive.ObjectIDFromHex(req.UserID)

	return Invitation{
		Name:        req.Name,
		Subdomain:   req.Subdomain,
		Published:   req.Published,
		TemplateIDs: []primitive.ObjectID{newTemplateID},
		UserIDs:     []primitive.ObjectID{newUserID},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (a Invitation) Validate() (err error) {
	return
}
