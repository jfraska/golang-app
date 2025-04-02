package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(router *gin.Engine, db *mongo.Database) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	r := router.Group("auth")
	{
		r.POST("register", handler.register)
		r.POST("login", handler.login)
		r.GET("logout", handler.logout)

		r.GET("google", handler.oauth)
		r.GET("google/callback", handler.oauthCallback)
	}
}
