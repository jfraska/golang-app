package infragin

import (
	"golang-app/infra/response"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

type Response struct {
	HttpCode int         `json:"-"`
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	Meta     interface{} `json:"meta,omitempty"`
	Error    *Error      `json:"error,omitempty"`
	Location string      `json:"location,omitempty"`
}

func NewResponse(params ...func(*Response) *Response) Response {
	var resp = Response{
		Success: true,
	}

	for _, param := range params {
		param(&resp)
	}

	return resp
}

func WithHttpCode(httpCode int) func(*Response) *Response {
	return func(r *Response) *Response {
		r.HttpCode = httpCode
		return r
	}
}

func WithMessage(message string) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Message = message
		return r
	}
}

func WithData(data interface{}) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Data = data
		return r
	}
}

func WithMeta(meta interface{}) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Meta = meta
		return r
	}
}

func WithLocation(location string) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Location = location
		return r
	}
}

func WithError(err error) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Success = false

		myErr, ok := err.(response.Error)
		if !ok {
			myErr = response.ErrorGeneral
		}

		if r.Error == nil {
			r.Error = &Error{}
		}

		r.Error.Message = myErr.Message
		r.Error.Code = myErr.Code
		r.HttpCode = myErr.HttpCode

		return r
	}
}

func (r Response) Send(ctx *gin.Context) {
	ctx.JSON(r.HttpCode, r)
}

func (r Response) Redirect(ctx *gin.Context) {
	ctx.Redirect(r.HttpCode, r.Location)
}
