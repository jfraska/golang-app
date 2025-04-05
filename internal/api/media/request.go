package media

import (
	"mime/multipart"

	"github.com/jfraska/golang-app/pkg/utils"
)

type CreateMediaPayload struct {
	InvitationID string                `form:"invitation_id" binding:"required"`
	Collection   string                `form:"collection" binding:"required"`
	File         *multipart.FileHeader `form:"file" binding:"required"`
}

type GetMediaPayload struct {
	utils.PaginationRequestPayload
	InvitationID string `form:"invitation_id" binding:"required"`
	Collection   string `form:"collection" binding:"required"`
	FileType     string `form:"file_type"`
}

type DeleteMediaPayload struct {
	ID string `uri:"id" binding:"required"`
}
