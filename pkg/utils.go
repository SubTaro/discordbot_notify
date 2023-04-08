package pkg

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Token string `json: Token`
	BotName string `json: BotName`
	ChannelID string `json: ChannelID`
}

func Read_conf(file_path string) *Config{
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var conf Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&conf); err != nil {
		log.Fatal(err)
	}

	return &conf
}
