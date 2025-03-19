package user

import (
	"fmt"
	infragin "golang-app/infra/gin"
	"golang-app/infra/response"
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

func (h handler) register(ctx *gin.Context) {
	var req RegisterRequestPayload

	if err := ctx.ShouldBindJSON(&req); err != nil {
		myErr := response.ErrorBadRequest
		infragin.NewResponse(
			infragin.WithMessage(err.Error()),
			infragin.WithError(myErr),
			infragin.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)

	}

	if err := h.svc.register(ctx, req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		infragin.NewResponse(
			infragin.WithMessage(err.Error()),
			infragin.WithError(myErr),
		).Send(ctx)
	}

	infragin.NewResponse(
		infragin.WithHttpCode(http.StatusCreated),
		infragin.WithMessage("register success"),
	).Send(ctx)

}

func (h handler) login(ctx *gin.Context) {
	var req LoginRequestPayload

	if err := ctx.ShouldBindJSON(&req); err != nil {
		myErr := response.ErrorBadRequest
		infragin.NewResponse(
			infragin.WithMessage(err.Error()),
			infragin.WithError(myErr),
			infragin.WithMessage("login fail"),
		).Send(ctx)
		return
	}

	token, err := h.svc.login(ctx, req)
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
		infragin.WithHttpCode(http.StatusCreated),
		infragin.WithData(map[string]interface{}{
			"access_token": token,
		}),
		infragin.WithMessage("login success"),
	).Send(ctx)
}

func (h handler) oauth(ctx *gin.Context) {
	var req OauthRequestPayload

	if err := ctx.ShouldBindQuery(&req); err != nil {
		myErr := response.ErrorBadRequest
		infragin.NewResponse(
			infragin.WithMessage(err.Error()),
			infragin.WithError(myErr),
			infragin.WithMessage("login fail"),
		).Send(ctx)
		return
	}

	url := h.svc.oauth(req)

	infragin.NewResponse(
		infragin.WithHttpCode(http.StatusTemporaryRedirect),
		infragin.WithData(map[string]interface{}{
			"url": url,
		}),
		infragin.WithMessage("temporary redirect url"),
	).Send(ctx)
}

func (h handler) oauthCallback(ctx *gin.Context) {
	var req OauthCallbackRequestPayload

	if err := ctx.ShouldBindQuery(&req); err != nil {
		myErr := response.ErrorBadRequest
		infragin.NewResponse(
			infragin.WithMessage(err.Error()),
			infragin.WithError(myErr),
			infragin.WithMessage("callback fail"),
		).Send(ctx)
		return
	}

	token, err := h.svc.oauthCallback(ctx, req)
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

	location := fmt.Sprintf("http://%s/api/auth/callback?access_token=%s", req.State, token)

	if req.State == "mobile" {
		return
	}

	infragin.NewResponse(
		infragin.WithHttpCode(http.StatusTemporaryRedirect),
		infragin.WithLocation(location),
	).Redirect(ctx)
}
