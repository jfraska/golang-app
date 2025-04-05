package invitation

import (
	"github.com/gin-gonic/gin"
	infragin "github.com/jfraska/golang-app/infra/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(router *gin.RouterGroup, db *mongo.Database) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	r := router.Group("invitation")
	{
		r.GET("/", infragin.CheckAuth(), handler.index)
		r.POST("/", infragin.CheckAuth(), handler.create)
	}

}
