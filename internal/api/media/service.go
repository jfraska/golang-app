package media

import (
	"context"
	"io"

	"github.com/jfraska/golang-app/infra/response"
	pkg "github.com/jfraska/golang-app/pkg/utils"
)

type Repository interface {
	CreateStorage(ctx context.Context, file []byte, model Media) (err error)
	CreateMedia(ctx context.Context, model Media) (err error)
	GetAllMediaByObjectName(ctx context.Context, model Media, pagination pkg.Pagination) ([]Media, pkg.Pagination, error)
	DeleteStorage(ctx context.Context, model Media) (err error)
	DeleteMedia(ctx context.Context, model Media) (err error)
	GetMediaByID(ctx context.Context, ID string) (model Media, err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) createMedia(ctx context.Context, req CreateMediaPayload) (err error) {
	media := NewMediaFromCreateMediaRequest(req)

	if err = media.Validate(); err != nil {
		return
	}

	file, err := req.File.Open()
	if err != nil {
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return
	}

	if err = s.repo.CreateStorage(ctx, fileBytes, media); err != nil {
		return
	}

	return s.repo.CreateMedia(ctx, media)

}

func (s service) listMedia(ctx context.Context, req GetMediaPayload) (media []Media, pagination pkg.Pagination, err error) {
	pagination = pkg.NewPaginationFromPaginationRequest(req.PaginationRequestPayload)
	model := NewMediaFromGetMediaRequest(req)

	media, pagination, err = s.repo.GetAllMediaByObjectName(ctx, model, pagination)
	if err != nil {
		if err == response.ErrNotFound {
			return
		}
		return

	}

	if len(media) == 0 {
		return
	}

	return
}

func (s service) deleteMedia(ctx context.Context, req DeleteMediaPayload) (err error) {

	model, err := s.repo.GetMediaByID(ctx, req.ID)
	if err != nil {
		return
	}

	if err = s.repo.DeleteMedia(ctx, model); err != nil {
		return
	}

	if err = s.repo.DeleteStorage(ctx, model); err != nil {
		return
	}

	return

}
