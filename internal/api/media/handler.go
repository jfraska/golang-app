package media

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
	var req CreateMediaPayload

	if err := ctx.ShouldBind(&req); err != nil {
		myErr := response.ErrorBadRequest
		infragin.NewResponse(
			infragin.WithMessage(err.Error()),
			infragin.WithError(myErr),
			infragin.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
		return
	}

	if err := h.svc.createMedia(ctx, req); err != nil {
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
		infragin.WithMessage("create Media success"),
	).Send(ctx)

}

func (h handler) index(ctx *gin.Context) {
	var req GetMediaPayload

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

	templates, pagination, err := h.svc.listMedia(ctx, req)
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
		infragin.WithData(templates),
		infragin.WithMeta(pagination),
	).Send(ctx)
}

func (h handler) delete(ctx *gin.Context) {
	var req DeleteMediaPayload

	if err := ctx.ShouldBindUri(&req); err != nil {
		myErr := response.ErrorBadRequest
		infragin.NewResponse(
			infragin.WithMessage(err.Error()),
			infragin.WithError(myErr),
			infragin.WithHttpCode(http.StatusBadRequest),
			infragin.WithMessage("invalid payload"),
		).Send(ctx)
		return
	}

	err := h.svc.deleteMedia(ctx, req)
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
		infragin.WithHttpCode(http.StatusNoContent),
		infragin.WithMessage("delete media success"),
	).Send(ctx)
}
