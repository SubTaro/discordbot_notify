package main

import (
	"fmt"
	"notify/bot"
	"notify/pkg"
	"os/exec"
	"strings"
	"flag"
)


func main() {
	flag.Parse()
	input_file := flag.Args()[0]
	commands := pkg.ReadCommand(input_file)
	stop_bot := make(chan bool)
	message := make(chan string)

	go bot.Start_bot("./conf/bot_conf.json", stop_bot, message)

	//fmt.Println(commands)

	for _, command := range commands {
		split_command := strings.Split(command, " ")

		if len(split_command) > 2 {
			output, err := exec.Command(split_command[0]).CombinedOutput()
			fmt.Printf("result: %s\nerror: %v\n\n", output, err)
			message<-"excuted"
		}else {
			output, err := exec.Command(split_command[0], split_command[1:]...).CombinedOutput()
			fmt.Printf("result: %s\nerror: %v\n\n", output, err)
			message<-"excuted"
		}
	}
	close(message)
	if <-stop_bot{
		fmt.Println("bot is stopped")
	}

	close(stop_bot)
	
	return
}
