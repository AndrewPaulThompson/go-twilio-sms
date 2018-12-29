package main

import (
    "log"

    "github.com/joho/godotenv"
)

// Main function
func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Setup up client
    client := Client{}
    client.setup()

    // Get message and send
    message := client.createMessage("number here", "HIII")
    client.send(message)
}
