package customize

import (
	"github.com/gin-gonic/gin"
	"github.com/jfraska/golang-app/infra/cache"
	infragin "github.com/jfraska/golang-app/infra/gin"
	"github.com/jfraska/golang-app/pkg/customize"
)

func Init(router *gin.RouterGroup, cache *cache.CacheMemory) {

	hub := customize.NewHub()
	handler := NewHandler(hub, cache)

	go hub.Run(cache)

	r := router.Group("customize")
	{
		r.GET("/", infragin.CheckAuth(), handler.GetRooms)
		r.GET("/:id/user", infragin.CheckAuth(), handler.GetClients)

		r.GET("/:id/ws", handler.JoinRoom)
	}
}
