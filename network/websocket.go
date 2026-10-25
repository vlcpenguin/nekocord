package network

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/vlcpenguin/nekocord/maps"
)

// Connect to discords gateway
func ConnectToDiscordGateway(serverAddress string, authToken string) {
	log.Println("[*] Connecting to discords gateway...")
	conn, _, err := websocket.DefaultDialer.Dial(serverAddress, nil)
	if err != nil {
		log.Fatal("Dial:", err)
	}
	defer conn.Close()

	log.Println("[+] Connected")
	log.Println("[*] Sending identify payload...")
	identifyPayload := maps.GetIdentifyPayload(authToken)

	// Send initial payload to identify yourself to discords gateway
	payloadBytes, err := json.Marshal(identifyPayload)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.WriteMessage(websocket.TextMessage, payloadBytes)
	if err != nil {
		log.Fatal("write:", err)
	}
	log.Println("[+] Sent identify payload")

	// Goroutine loop for reading websocket messages
	go func() {
		err := read(conn)
		if err != nil {
			log.Println("gateway read: %v", err)
		}
	}()

	select {}
}
