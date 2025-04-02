package invitation

import "context"

type Repository interface {
	CreateInvitation(ctx context.Context, model Invitation) error
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
