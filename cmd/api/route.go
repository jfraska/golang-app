package main

import (
	"github.com/jfraska/golang-app/infra/cache"
	"github.com/jfraska/golang-app/internal/api/invitation"
	"github.com/jfraska/golang-app/internal/api/media"
	"github.com/jfraska/golang-app/internal/api/template"
	"github.com/jfraska/golang-app/internal/api/user"
	"github.com/jfraska/golang-app/internal/websocket/customize"
	"github.com/minio/minio-go/v7"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func initRoute(router *gin.Engine, db *mongo.Database, sdb *minio.Client, cache *cache.CacheMemory) {
	user.Init(router, db)

	v1 := router.Group("v1")
	invitation.Init(v1, db)
	template.Init(v1, db)
	media.Init(v1, db, sdb)
	customize.Init(v1, cache)
}
