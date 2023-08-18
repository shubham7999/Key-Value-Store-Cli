package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var keyValueStore = make(map[string]string)

func main() {
	fmt.Println("Advanced Key-Value Store")
	fmt.Println("------------------------")

	loadData()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		command := strings.Fields(input)

		if len(command) == 0 {
			continue
		}

		switch command[0] {
		case "GET":
			if len(command) != 2 {
				fmt.Println("Invalid number of arguments for GET command.")
				continue
			}
			key := command[1]
			value, ok := keyValueStore[key]
			if !ok {
				fmt.Println("Key not found:", key)
			} else {
				fmt.Printf("Value for key '%v' is '%v'\n", key , value)
			}

		case "SET":
			if len(command) != 3 {
				fmt.Println("Invalid number of arguments for SET command.")
				continue
			}
			key := command[1]
			value := command[2]
			keyValueStore[key] = value
			fmt.Printf("Store key '%v' with value '%v'\n", key , value)

		case "DELETE":
			if len(command) != 2 {
				fmt.Println("Invalid number of arguments for DELETE command.")
				continue
			}
			key := command[1]
			_, ok := keyValueStore[key]
			if !ok {
				fmt.Println("Key not found:", key)
			} else {
				delete(keyValueStore, key)
				fmt.Printf("Key '%v' deleted.\n", key)
			}

		case "SAVE":
			saveData()
			fmt.Println("Data saved to file.")

		case "LOAD":
			loadData()
			fmt.Println("Data loaded from file.")

		case "EXIT":
			saveData()
			fmt.Println("Exiting the program.")
			os.Exit(0)

		case "LIST":
			fmt.Println("Stored keys:")
			for key := range keyValueStore {
				fmt.Printf("'%v' \n", key)
			}

		default:
			fmt.Println("Invalid command:", command[0])
		}
	}
}

func saveData() {
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for key, value := range keyValueStore {
		_, _ = fmt.Fprintf(file, "%s=%s\n", key, value)
	}
}

func loadData() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("No existing data file found.")
		return
	}
	defer file.Close()

	keyValueStore = make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			keyValueStore[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading data file:", err)
	}
}

