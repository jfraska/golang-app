package customize

import (
	"context"

	"github.com/jfraska/golang-app/infra/cache"
)

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run(cache *cache.CacheMemory) {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
				} else {
					h.Broadcast <- &Message{
						RoomID:   cl.RoomID,
						Username: cl.Username,
					}
				}

			} else {
				h.Rooms[cl.RoomID] = &Room{
					ID:      cl.RoomID,
					Name:    cl.RoomID,
					Clients: make(map[string]*Client),
				}

				r := h.Rooms[cl.RoomID]
				r.Clients[cl.ID] = cl
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {

					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)

					if len(h.Rooms[cl.RoomID].Clients) == 0 {
						// delete in memory
						delete(h.Rooms, cl.RoomID)

						// delete in redis database
						cache.Del(context.Background(), cl.RoomID)
					}

				}
			}

		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {

				for _, cl := range h.Rooms[m.RoomID].Clients {
					cl.Message <- m
				}
			}
		}
	}
}
