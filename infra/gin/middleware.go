package infragin

import (
	"golang-app/infra/response"
	"golang-app/infra/session"
	"golang-app/internal/config"
	pkg "golang-app/pkg/utils"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		if authorization == "" {
			NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(ctx)
			ctx.Abort()
		}

		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			log.Println("Token invalid")
			NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(ctx)
			ctx.Abort()
		}

		token := bearer[1]

		publicId, role, err := pkg.ValidateToken(token, config.Cfg.Encryption.JWTSecret)
		if err != nil {
			log.Println(err.Error())
			NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(ctx)
			ctx.Abort()
		}

		_, err = session.Store.Get(ctx, publicId)
		if err != nil {
			NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(ctx)
			ctx.Abort()
		}

		ctx.Set("ROLE", role)
		ctx.Set("PUBLIC_ID", publicId)

		ctx.Next()
	}
}

func CheckRoles(authorizedRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("ROLE")
		if role == "" {
			NewResponse(
				WithError(response.ErrorForbiddenAccess),
			).Send(ctx)
			ctx.Abort()
		}

		isExists := false
		for _, authorizedRole := range authorizedRoles {
			if role == authorizedRole {
				isExists = true
				break
			}
		}

		if !isExists {
			NewResponse(
				WithError(response.ErrorForbiddenAccess),
			).Send(ctx)
			ctx.Abort()
		}

		ctx.Next()
	}
}
