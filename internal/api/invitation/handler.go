package invitation

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	infragin "github.com/jfraska/golang-app/infra/gin"
	"github.com/jfraska/golang-app/infra/response"
	"github.com/jfraska/golang-app/pkg/utils"
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

	publicID, _ := ctx.Get("PUBLIC_ID")
	req.UserID = fmt.Sprintf("%v", publicID)

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

func (h handler) index(ctx *gin.Context) {
	var req utils.PaginationRequestPayload

	if err := ctx.ShouldBindQuery(&req); err != nil {
		myErr := response.ErrorBadRequest
		infragin.NewResponse(
			infragin.WithMessage(err.Error()),
			infragin.WithError(myErr),
			infragin.WithHttpCode(http.StatusBadRequest),
			infragin.WithMessage("invalid payload"),
		).Send(ctx)
		return
	}

	publicID, _ := ctx.Get("PUBLIC_ID")

	invitations, pagination, err := h.svc.listInvitations(ctx, fmt.Sprintf("%v", publicID), req)
	if err != nil {
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
		infragin.WithHttpCode(http.StatusOK),
		infragin.WithMessage("get list media success"),
		infragin.WithData(invitations),
		infragin.WithMeta(pagination),
	).Send(ctx)
}
