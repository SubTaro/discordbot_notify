package main

import (
	"fmt"
	"notify/bot"
	"notify/pkg"
	"os/exec"
	"strings"
)

func main() {
	commands := pkg.ReadCommand("input.ntfy")
	stop_bot := make(chan bool)
	message := make(chan string)

	go bot.Start_bot("./conf/conf.json", stop_bot, message)

	//fmt.Println(commands)

	for _, command := range commands {
		split_command := strings.Split(command, " ")

		if len(split_command) > 2 {
			output, err := exec.Command(split_command[0]).CombinedOutput()
			fmt.Printf("result: %s\nerror: %v\n\n", output, err)
			message<-string(output)
		}else {
			output, err := exec.Command(split_command[0], split_command[1:]...).CombinedOutput()
			fmt.Printf("result: %s\nerror: %v\n\n", output, err)
			message<-string(output)
		}
	}
	stop_bot<-true

	close(message)
	close(stop_bot)
	
	return
}
