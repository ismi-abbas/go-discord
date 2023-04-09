package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

// own datatype
type configStruct struct {
	Token     string `json:"token"`
	BotPrefix string `json:"bot_prefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	config = &configStruct{}

	file, err := os.ReadFile("config.json")

	if err != nil {
		fmt.Println("Error reading config file: ", err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println("Error unmarshalling config file: ", err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
