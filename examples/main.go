package main

import (
	"github.com/vlcpenguin/nekocord"
	"log"
	"os"
)

var discord_token string = os.Getenv("discord_token")

func main() {
	nekocord.SetCallback("MESSAGE_CREATE", func(m map[string]interface{}) {
		content, ok := m["content"]
		if !ok {
			log.Println("content is not valid type or doesnt exist")
			return
		}
		authorObject, ok := m["author"].(map[string]interface{})
		if !ok {
			log.Println("author nil")
			return
		}
		log.Printf("%s //=> (%s)", content, authorObject["username"])
	})
	nekocord.StartWebsocketSession(discord_token)
}
