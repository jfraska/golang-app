package media

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Media struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Filename   string             `bson:"filename" json:"filename"`
	FileType   string             `bson:"file_type" json:"file_type"`
	FileSize   int64              `bson:"file_size" json:"file_size"`
	Collection string             `bson:"collection" json:"collection"`
	CreatedAt  time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`

	InvitationID primitive.ObjectID `bson:"invitation_id" json:"invitation_id"`
}

func NewMediaFromCreateMediaRequest(req CreateMediaPayload) Media {

	newID, _ := primitive.ObjectIDFromHex(req.InvitationID)
	newFilename := fmt.Sprintf("%s_%s%s", strings.TrimSuffix(req.File.Filename, filepath.Ext(req.File.Filename)), time.Now().Format("20060102150405"), filepath.Ext(req.File.Filename))

	return Media{
		InvitationID: newID,
		Filename:     newFilename,
		FileType:     req.File.Header.Get("Content-Type"),
		FileSize:     req.File.Size,
		Collection:   req.Collection,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func NewMediaFromGetMediaRequest(req GetMediaPayload) Media {

	newID, _ := primitive.ObjectIDFromHex(req.InvitationID)

	return Media{
		InvitationID: newID,
		FileType:     req.FileType,
		Collection:   req.Collection,
	}
}

func (a Media) Validate() (err error) {
	return
}
