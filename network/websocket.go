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

	// Reading websocket messages that pass through the websocket connection
	done := make(chan error, 1)

	go func() {
		done <- read(conn)
	}()
	if err := <-done; err != nil {
		log.Printf("gateway error %v", err)
	}
	log.Println("[*] Discord gateway connection closed")
}
