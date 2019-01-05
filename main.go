package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// NumberTo is the environment variable name where therecipient phone number should be read from
const NumberTo string = "NUMBER_TO"

// Main function
func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup up client
	client := NewClient()

	// Get message and send
	message := client.createMessage(os.Getenv(NumberTo), "HIII")
	client.send(message)
}
