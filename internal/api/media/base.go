package media

import (
	"github.com/gin-gonic/gin"
	infragin "github.com/jfraska/golang-app/infra/gin"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(router *gin.RouterGroup, db *mongo.Database, sdb *minio.Client) {
	repo := newRepository(db, sdb)
	svc := newService(repo)
	handler := newHandler(svc)

	r := router.Group("media")
	{
		r.GET("/", infragin.CheckAuth(), handler.index)
		r.POST("/", infragin.CheckAuth(), handler.create)
		r.DELETE("/:id", infragin.CheckAuth(), handler.delete)
	}
}
