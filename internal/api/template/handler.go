package template

import (
	"net/http"

	infragin "github.com/jfraska/golang-app/infra/gin"
	"github.com/jfraska/golang-app/infra/response"
	pkg "github.com/jfraska/golang-app/pkg/utils"

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

func (h handler) create(ctx *gin.Context) {
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

func (h handler) index(ctx *gin.Context) {
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

func (h handler) show(ctx *gin.Context) {
	ID := ctx.Param("id")

	template, err := h.svc.getTemplate(ctx, ID)
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

func (h handler) delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	if err := h.svc.deleteTemplate(ctx, ID); err != nil {
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

// func (h handler) Listen() {
// 	ctx := context.Background()

// 	h.broker.Subscribe(ctx, "updateTemplate", func(message broker.Message) {
// 		fmt.Println("Event updateTemplate diterima:", message)

// 	})

// 	h.broker.Subscribe(ctx, "getTemplate", func(message broker.Message) {
// 		fmt.Println("Event getTemplate diterima:", message)
// 	})
// }
