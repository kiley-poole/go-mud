package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var commandMap = make(map[string]func(string))

//Main stuff here y'all
func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	populateCommandMap()
	if err := inputLoop(); err != nil {
		log.Fatalf("%v", err)
	}
}

//Input loop for processing user commands
func inputLoop() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please Enter Your Command: ")
	for scanner.Scan() {
		processCommand(scanner.Text())
		fmt.Print("Enter another command: ")
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

//Processes command passed in by user
func processCommand(command string) error {
	words := strings.Fields(command)
	if len(words) == 0 {
		return nil
	}
	if f, exists := commandMap[strings.ToLower(words[0])]; exists {
		f(command)
	} else {
		fmt.Println("Invalid command.")
	}
	return nil
}

//Populates commandMap map with partial and full string name as key and the function as the value
func addCommand(command string, commandFunc func(string)) {
	for i := range command {
		if i == 0 {
			continue
		}
		prefix := command[:i]
		commandMap[prefix] = commandFunc
	}
	commandMap[command] = commandFunc
}

func populateCommandMap() {
	addCommand("up", lookUp)
	addCommand("down", lookDown)
	addCommand("north", goNorth)
	addCommand("south", goSouth)
	addCommand("east", goEast)
	addCommand("west", goWest)
}

func lookUp(s string) {
	fmt.Println("You look up!")
}

func lookDown(s string) {
	fmt.Println("You look down")
}

func goNorth(s string) {
	fmt.Println("You head north")
}

func goSouth(s string) {
	fmt.Println("You head south")
}

func goEast(s string) {
	fmt.Println("You head east")
}

func goWest(s string) {
	fmt.Println("You head west")
}
