package main

import (
	"fmt"

	"github.com/ismi-abbas/discord/bot"
	"github.com/ismi-abbas/discord/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
}
