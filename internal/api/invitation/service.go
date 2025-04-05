package invitation

import (
	"context"

	"github.com/jfraska/golang-app/infra/response"
	"github.com/jfraska/golang-app/pkg/utils"
)

type Repository interface {
	CreateInvitation(ctx context.Context, model Invitation) error
	GetAllInvitations(ctx context.Context, UserID string, pagination utils.Pagination) ([]Invitation, utils.Pagination, error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) createInvitation(ctx context.Context, req CreateInvitationRequestPayload) (err error) {
	template := NewInvitationFromCreateInvitationRequest(req)

	if err = template.Validate(); err != nil {
		return
	}

	return s.repo.CreateInvitation(ctx, template)

}

func (s service) listInvitations(ctx context.Context, UserID string, req utils.PaginationRequestPayload) (invitation []Invitation, pagination utils.Pagination, err error) {
	pagination = utils.NewPaginationFromPaginationRequest(req)

	invitation, pagination, err = s.repo.GetAllInvitations(ctx, UserID, pagination)
	if err != nil {
		if err == response.ErrNotFound {
			return
		}
		return

	}

	if len(invitation) == 0 {
		invitation = []Invitation{}
		return
	}

	return
}
