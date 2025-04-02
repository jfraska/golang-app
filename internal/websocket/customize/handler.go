package customize

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jfraska/golang-app/infra/cache"
	infragin "github.com/jfraska/golang-app/infra/gin"
	"github.com/jfraska/golang-app/infra/response"
	"github.com/jfraska/golang-app/pkg/customize"
)

type Handler struct {
	hub   *customize.Hub
	cache *cache.CacheMemory
}

func NewHandler(hub *customize.Hub, cache *cache.CacheMemory) *Handler {
	return &Handler{
		hub:   hub,
		cache: cache,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h Handler) JoinRoom(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		infragin.NewResponse(
			infragin.WithMessage(err.Error()),
			infragin.WithError(myErr),
		).Send(ctx)
		return
	}

	roomID := ctx.Param("id")
	clientID := ctx.Query("userId")
	username := ctx.Query("username")

	defer conn.Close()

	body, err := h.cache.Get(ctx, roomID)
	if err != nil {
		body = cache.CacheStore{}

		if err = h.cache.Set(ctx, roomID, body); err != nil {
			fmt.Println("error set cache:", err)
			return
		}
	}

	m := &customize.Message{
		Content:  body.Content,
		RoomID:   roomID,
		Username: username,
	}

	cl := &customize.Client{
		Conn:     conn,
		Message:  make(chan *customize.Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage(h.cache)
	cl.ReadMessage(h.hub)
}

func (h Handler) GetRooms(ctx *gin.Context) {
	rooms := make([]RoomResponse, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomResponse{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	infragin.NewResponse(
		infragin.WithHttpCode(http.StatusOK),
		infragin.WithMessage("get list rooms success"),
		infragin.WithData(rooms),
	).Send(ctx)
}

func (h Handler) GetClients(ctx *gin.Context) {
	var clients []ClientResponse
	roomId := ctx.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientResponse, 0)
		infragin.NewResponse(
			infragin.WithHttpCode(http.StatusOK),
			infragin.WithMessage("get list clients success"),
			infragin.WithData(clients),
		).Send(ctx)
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientResponse{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	infragin.NewResponse(
		infragin.WithHttpCode(http.StatusOK),
		infragin.WithMessage("get list clients success"),
		infragin.WithData(clients),
	).Send(ctx)
}
