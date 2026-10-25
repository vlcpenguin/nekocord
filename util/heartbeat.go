package util

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vlcpenguin/nekocord/models"
)

var globSeqNum *int = nil

func sendHeartBeat(conn *websocket.Conn, sequenceNumber *int) error {
	log.Printf("[*] Sending heartbeat... %v", ValueOrNil(sequenceNumber))
	payload := structs.HeartbeatPayload{
		Op: 1,
		D:  sequenceNumber,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal heartbeat: %w", err)
	}
	err = conn.WriteMessage(websocket.TextMessage, payloadBytes)
	if err != nil {
		return fmt.Errorf("send heartbeat: %w", err)
	}
	log.Println("[+] Sent heartbeat")
	return nil
}

// Send heartbeat between each heartbeat interval
func startHeartBeat(conn *websocket.Conn, interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		defer ticker.Stop()
		send := func() {
			if err := sendHeartBeat(conn, globSeqNum); err != nil {
				log.Printf("[-] Heartbeat failed: %v", err)
			}
		}
		send()

		for range ticker.C {
			send()
		}
	}()
}

func handleHeartbeat(msg []byte, conn *websocket.Conn) error {
	var Message structs.DiscordGatewayMessage

	err := json.Unmarshal(msg, &Message)
	if err != nil {
		return fmt.Errorf("unmarshal heartbeat message %v: ", err)
	}

	if Message.SequenceNumber != 0 {
		if globSeqNum == nil {
			globSeqNum = new(int)
		}
		*globSeqNum = Message.SequenceNumber
	}

	switch Message.GatewayOpcode {

	// Opcode: Hello (Returns a heartbeat interval)
	case 10:
		dataMap, ok := Message.EventData.(map[string]interface{})
		if !ok {
			return nil
		}
		heartbeatInterval, ok := dataMap["heartbeat_interval"].(float64)
		if !ok {
			log.Println("No heartbeat_interval found")
			break
		}

		startHeartBeat(conn, time.Duration(heartbeatInterval)*time.Millisecond)

	// Opcode: OK (Tells the client that the heartbeat payload was good)
	case 11:
		log.Println("[+] Heartbeat Acknowledged")

	}
	return nil
}
