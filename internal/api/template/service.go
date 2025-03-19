package template

import (
	"context"
	"golang-app/infra/response"
	pkg "golang-app/pkg/utils"
)

type Repository interface {
	CreateTemplate(ctx context.Context, model Template) error
	GetTemplateBySlug(ctx context.Context, slug string) (model Template, err error)
	GetAllTemplates(ctx context.Context, model pkg.Pagination) ([]Template, pkg.Pagination, error)
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

func (s service) TemplateDetail(ctx context.Context, req GetTemplateRequestPayload) (model Template, err error) {
	model, err = s.repo.GetTemplateBySlug(ctx, req.Slug)
	if err != nil {
		return
	}
	return
}
