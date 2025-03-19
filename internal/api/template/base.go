package template

import (
	infragin "golang-app/infra/gin"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(router *gin.RouterGroup, db *mongo.Database) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	r := router.Group("template")
	{
		r.GET("/", handler.getListTemplates)
		r.GET("/:slug", infragin.CheckAuth(), handler.GetTemplateDetail)
		r.POST("/", infragin.CheckAuth(), handler.createTemplate)
	}
}
