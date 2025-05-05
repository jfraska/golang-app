package template

import (
	"github.com/gin-gonic/gin"
	infragin "github.com/jfraska/golang-app/infra/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(router *gin.RouterGroup, db *mongo.Database) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	// go handler.Listen()

	r := router.Group("template")
	{
		r.GET("/", handler.index)
		r.GET("/:id", handler.show)
		r.POST("/", infragin.CheckAuth(), handler.create)
		r.DELETE("/:id", infragin.CheckAuth(), handler.delete)
	}

}
