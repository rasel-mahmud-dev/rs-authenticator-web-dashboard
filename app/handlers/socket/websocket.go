package socket

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading csonnecjjjjjtinon:", err)
		return
	}

	fmt.Println("WebSocket cossjhhkkjjkdsdsdsasasasnnection opened:", conn.RemoteAddr())

	clients[conn] = true
	defer func() {
		fmt.Println("WebSocket connection closed:", conn.RemoteAddr())
		delete(clients, conn)
		conn.Close()
	}()

	err = conn.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(10*time.Second))
	if err != nil {
		fmt.Println("Error sending ping message:", err)
		return
	}

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket read error (closing connection):", err)
			break
		}
	}
}
