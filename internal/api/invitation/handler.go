package invitation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	infragin "github.com/jfraska/golang-app/infra/gin"
	"github.com/jfraska/golang-app/infra/response"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) create(ctx *gin.Context) {
	var req CreateInvitationRequestPayload

	if err := ctx.ShouldBindJSON(&req); err != nil {
		myErr := response.ErrorBadRequest
		infragin.NewResponse(
			infragin.WithMessage(err.Error()),
			infragin.WithError(myErr),
			infragin.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
		return
	}

	if err := h.svc.createInvitation(ctx, req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		infragin.NewResponse(
			infragin.WithMessage(err.Error()),
			infragin.WithError(myErr),
		).Send(ctx)
		return
	}

	infragin.NewResponse(
		infragin.WithHttpCode(http.StatusCreated),
		infragin.WithMessage("create invitation success"),
	).Send(ctx)

}
