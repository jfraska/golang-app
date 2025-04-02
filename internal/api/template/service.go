package template

import (
	"context"

	"github.com/jfraska/golang-app/infra/response"
	pkg "github.com/jfraska/golang-app/pkg/utils"
)

type Repository interface {
	CreateTemplate(ctx context.Context, model Template) error
	GetTemplateByID(ctx context.Context, ID string) (model Template, err error)
	GetAllTemplates(ctx context.Context, model pkg.Pagination) ([]Template, pkg.Pagination, error)
	UpdateTemplateByID(ctx context.Context, model Template) (err error)
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
			return []Template{}, pkg.Pagination{}, nil
		}
		return

	}

	if len(templates) == 0 {
		return []Template{}, pkg.Pagination{}, nil
	}

	return
}

func (s service) TemplateDetail(ctx context.Context, ID string) (model Template, err error) {
	model, err = s.repo.GetTemplateByID(ctx, ID)
	if err != nil {
		return
	}
	return
}

func (s service) UpdateTemplateByCustomize(ctx context.Context, model Template) (err error) {
	return s.repo.UpdateTemplateByID(ctx, model)
}

func (s service) GetTemplate(ctx context.Context, ID string) (model Template, err error) {
	model, err = s.repo.GetTemplateByID(ctx, ID)
	if err != nil {
		return
	}
	return
}
