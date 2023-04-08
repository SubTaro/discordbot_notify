package bot

import (
	"fmt"
	"log"
	"notify/pkg"

	"github.com/bwmarrin/discordgo"
)

func Start_bot(config_file string, stop_bot chan bool, message chan string) {
	//Configの読み込み
	conf := pkg.Read_conf(config_file)
	//botのスタート
	bot, err := discordgo.New(conf.BotName)
	bot.Token = conf.Token

	if err != nil {
		fmt.Println("Error login")
		log.Fatal(err)
	}

	defer bot.Close()
	fmt.Println("Listening...")

	for {
		msg, ok := <-message
		if ok != true {
			break
		}
		bot.ChannelMessageSend(conf.ChannelID, msg)
	}

	if <-stop_bot{
		return
	}
}

func Send_message(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)

	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Sending Message Error: ", err)
	}
}
