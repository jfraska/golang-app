package template

import (
	"context"

	"github.com/jfraska/golang-app/infra/response"
	pkg "github.com/jfraska/golang-app/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	CreateTemplate(ctx context.Context, model Template) error
	GetTemplateByID(ctx context.Context, ID primitive.ObjectID) (model Template, err error)
	GetAllTemplates(ctx context.Context, model pkg.Pagination) ([]Template, pkg.Pagination, error)
	UpdateTemplateByID(ctx context.Context, model Template) (err error)
	DeleteTemplate(ctx context.Context, ID primitive.ObjectID) (err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) createTemplate(ctx context.Context, req CreateTemplateRequestPayload) (err error) {
	template := NewTemplateFromCreateTemplateRequest(req)

	if err = template.Validate(); err != nil {
		return
	}

	return s.repo.CreateTemplate(ctx, template)

}

func (s service) listTemplates(ctx context.Context, req pkg.PaginationRequestPayload) (templates []Template, pagination pkg.Pagination, err error) {
	pagination = pkg.NewPaginationFromPaginationRequest(req)

	templates, pagination, err = s.repo.GetAllTemplates(ctx, pagination)
	if err != nil {
		if err == response.ErrNotFound {
			return
		}
		return

	}

	return
}

func (s service) getTemplate(ctx context.Context, ID string) (model Template, err error) {
	newID, _ := pkg.ConvertObjectID(ID)

	model, err = s.repo.GetTemplateByID(ctx, newID)
	if err != nil {
		return
	}
	return
}

func (s service) updateTemplateByCustomize(ctx context.Context, model Template) (err error) {
	return s.repo.UpdateTemplateByID(ctx, model)
}

func (s service) deleteTemplate(ctx context.Context, ID string) (err error) {

	newID, _ := pkg.ConvertObjectID(ID)

	err = s.repo.DeleteTemplate(ctx, newID)
	if err != nil {
		return
	}
	return

}
