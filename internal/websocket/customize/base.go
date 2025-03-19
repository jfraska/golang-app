package customize

import (
	pkg "golang-app/pkg/ws"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {

	hub := pkg.NewHub()
	handler := NewHandler(hub)

	go hub.Run()

	r := router.Group("customize")
	{
		r.POST("/", handler.CreateRoom)
		r.GET("/", handler.GetRooms)
		r.GET("/:id/user", handler.GetClients)

		r.GET("/:id/ws", handler.JoinRoom)
	}
}
