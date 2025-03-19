package main

import (
	"golang-app/internal/api/template"
	"golang-app/internal/api/user"
	"golang-app/internal/websocket/customize"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func initRoute(router *gin.Engine, db *mongo.Database) {
	user.Init(router, db)

	v1 := router.Group("v1")
	template.Init(v1, db)
	customize.Init(v1)
}
