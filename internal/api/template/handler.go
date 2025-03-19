package template

import (
	infragin "golang-app/infra/gin"
	"golang-app/infra/response"
	pkg "golang-app/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) createTemplate(ctx *gin.Context) {
	var req CreateTemplateRequestPayload

	if err := ctx.ShouldBindJSON(&req); err != nil {
		myErr := response.ErrorBadRequest
		infragin.NewResponse(
			infragin.WithMessage(err.Error()),
			infragin.WithError(myErr),
			infragin.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
		return
	}

	if err := h.svc.createTemplate(ctx, req); err != nil {
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
		infragin.WithMessage("create template success"),
	).Send(ctx)

}

func (h handler) getListTemplates(ctx *gin.Context) {
	var req pkg.PaginationRequestPayload

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

	templates, pagination, err := h.svc.listTemplates(ctx, req)
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

	templateList := NewTemplateListResponseFromEntity(templates)

	infragin.NewResponse(
		infragin.WithHttpCode(http.StatusOK),
		infragin.WithMessage("get list templates success"),
		infragin.WithData(templateList),
		infragin.WithMeta(pagination),
	).Send(ctx)
}

func (h handler) GetTemplateDetail(ctx *gin.Context) {
	var req GetTemplateRequestPayload

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

	template, err := h.svc.TemplateDetail(ctx, req)
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

	templateDetail := NewTemplateDetailResponseFromEntity(template)

	infragin.NewResponse(
		infragin.WithHttpCode(http.StatusOK),
		infragin.WithMessage("get template detail success"),
		infragin.WithData(templateDetail),
	).Send(ctx)
}
