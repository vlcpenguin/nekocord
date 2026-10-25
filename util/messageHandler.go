package util

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	structs "github.com/vlcpenguin/nekocord/models"
)

func HandleMessage(Message []byte, conn *websocket.Conn) error {
	// Room for improvement i didn't know you could just turn a json into an interface with no
	// useless structs.
	//
	// bug prone...
	var msg structs.DiscordGatewayMessage
	err := json.Unmarshal(Message, &msg)
	if err != nil {
		return fmt.Errorf("message unmarshal: %v", err)
	}

	// Heartbeat
	err = handleHeartbeat(Message, conn)
	if err != nil {
		return fmt.Errorf("handle heartbeat: %v", err)
	}

	eData, Ok := msg.EventData.(map[string]interface{})
	if !Ok {
		return nil
	}

	// Event Handler
	if callback, ok := Callbacks[msg.EventName]; ok {
		callback(eData)
	}
	return nil
}
