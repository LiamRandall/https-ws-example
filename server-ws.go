// You need to install the Gorilla WebSocket package
// Run: go get -u github.com/gorilla/websocket
package main

import (
	"fmt"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        // Here you can check the origin and reject requests from unexpected domains
        return true
    },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error upgrading to WebSocket:", err)
        return
    }
    defer conn.Close()

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Error reading message:", err)
            break
        }
        fmt.Printf("Received message: %s\n", p)
        if err = conn.WriteMessage(messageType, p); err != nil {
            fmt.Println("Error writing message:", err)
            break
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleWebSocket)

    fmt.Println("WebSocket server listening on ws://localhost:8081")
    http.ListenAndServe(":8081", nil)
}
