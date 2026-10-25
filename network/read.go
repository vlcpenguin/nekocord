package network

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/vlcpenguin/nekocord/util"
)

func read(conn *websocket.Conn) error {
	for {
		_, Message, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("read message: %v", err)
		}
		// If we get a message then send it to the message handler
		if Message != nil {
			err := util.HandleMessage(Message, conn)
			if err != nil {
				return fmt.Errorf("handle message: %v", err)
			}
		}
	}
}
