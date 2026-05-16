package util

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/vlcpenguin/nekocord/models"
	"log"
)

func MessageHandler(Message []byte, conn websocket.Conn) {
	// Room for improvement i didn't know you could just turn a json into an interface with no
	// useless structs.
	//
	// bug prone...
	var msg structs.DiscordGatewayMessage
	err := json.Unmarshal(Message, &msg)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Heartbeat
	handleHeartbeat(Message, conn)

	eData, Ok := msg.EventData.(map[string]interface{})
	if !Ok {
		return
	}

	// Event Handler
	if callback, ok := Callbacks[msg.EventName]; ok {
		callback(eData)
	}

}
