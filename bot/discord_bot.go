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

	fmt.Println("Listening...")

	for {
		msg, ok := <-message
		if ok != true {
			break
		}
		bot.ChannelMessageSend(conf.ChannelID, msg)
	}
	bot.ChannelMessageSend(conf.ChannelID, "End")
	stop_bot<-true
	bot.Close()

	return
}
