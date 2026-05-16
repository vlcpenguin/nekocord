package nekocord

import (
	"github.com/vlcpenguin/nekocord/network"
	"github.com/vlcpenguin/nekocord/util"
)

func SetCallback(eventName string, callback func(map[string]interface{})) {
	util.SetCallback(eventName, callback)
}

func StartWebsocketSession(auth_token string) {
	network.ConnectToDiscordGateway("wss://gateway.discord.gg/?encoding=json&v=9", auth_token)
}
