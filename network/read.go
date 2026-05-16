package network

import (
	"github.com/gorilla/websocket"
	"github.com/vlcpenguin/nekocord/util"
	"log"
)

func read(conn websocket.Conn) {
	for {
		_, Message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("Error:", err)
			return
		}
		// If the message isn't nil ( Exists ) then send it to the messageHandler
		if Message != nil {
			util.MessageHandler(Message, conn)
		}
	}
}
