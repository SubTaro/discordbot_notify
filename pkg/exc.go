package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func ReadCommand(input string) []string{
	commands := []string{}
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	return commands
}
