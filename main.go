package main

import (
	cf "github.com/ataka/configFile"
	cw "github.com/ataka/go-chatwork"
)

type Config struct {
	ChatWorkConfig `json:"chatwork"`
}

type ChatWorkConfig struct {
	ChatWorkToken string `json:"token"`
	RoomId        int64  `json:"roomId"`
}

func main() {
	var config Config
	configFile := cf.NewConfigFile(".monch", &config)
	configFile.Read()
	
	chatwork := cw.NewChatwork(config.ChatWorkToken)
	roomId := config.RoomId
	message := cw.NewMessage(roomId, "テストだよ")
	chatwork.CreateMessage(message)
}
