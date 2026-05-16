package util

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/vlcpenguin/nekocord/models"
	"log"
	"time"
)

var globSeqNum *int = nil

func sendHeartBeat(conn websocket.Conn, sequenceNumber *int) {
	if sequenceNumber != nil {
		log.Println("[*] Sending heartbeat...", *sequenceNumber)
	} else {
		log.Println("[*] Sending heartbeat... <nil>")
	}
	payload := structs.HeartbeatPayload{
		Op: 1,
		D:  sequenceNumber,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, payloadBytes)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("[+] Sent heartbeat")
	log.Println("[*] Waiting for server Acknowledgement...")
}

// Send heartbeat inbetween given heartbeat interval
func startHeartBeat(conn websocket.Conn, heartbeatInterval int) {
	ticker := time.NewTicker(time.Duration(heartbeatInterval) * time.Millisecond)

	go func() {
		sendHeartBeat(conn, globSeqNum)

		for range ticker.C {
			sendHeartBeat(conn, globSeqNum)
		}
	}()

}

func handleHeartbeat(msg []byte, conn websocket.Conn) {
	var Message structs.DiscordGatewayMessage

	err := json.Unmarshal(msg, &Message)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Get Sequence Number \\
	if Message.SequenceNumber != 0 {
		if globSeqNum == nil {
			globSeqNum = new(int)
		}
		*globSeqNum = Message.SequenceNumber
	}

	// Opcode Logic \\
	switch Message.GatewayOpcode {

	// Opcode 10 Hello

	case 10:
		dataMap, ok := Message.EventData.(map[string]interface{})
		if !ok {
			return
		}
		heartbeatInterval, ok := dataMap["heartbeat_interval"].(float64)
		if !ok {
			log.Println("No heartbeat_interval found")
			break
		}

		startHeartBeat(conn, int(heartbeatInterval))

	// OK (Heartbeat Acknowledged)
	case 11:
		log.Println("[+] Heartbeat Acknowledged")

	}
}
