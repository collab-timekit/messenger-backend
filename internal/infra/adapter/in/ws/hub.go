package ws

// Hub manages the registration and unregistration of WebSocket clients.
type Hub struct {
    clients    map[string]*Client
    register   chan *Client
    unregister chan *Client
}

// NewHub creates and returns a new instance of Hub.
func NewHub() *Hub {
    return &Hub{
        clients:    make(map[string]*Client),
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
}

// Run starts the Hub's main loop to handle client registration and unregistration.
func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            if old, ok := h.clients[client.userID]; ok {
                old.conn.Close()
            }
            h.clients[client.userID] = client

        case client := <-h.unregister:
            if _, ok := h.clients[client.userID]; ok {
                delete(h.clients, client.userID)
                close(client.send)
            }
        }
    }
}

// GetClient retrieves a client by their userID from the Hub.
func (h *Hub) GetClient(userID string) (*Client, bool) {
    client, ok := h.clients[userID]
    return client, ok
}