package customize

type ClientResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type RoomResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
