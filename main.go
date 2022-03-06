package main

import (
	"gochat/chat"
	"log"
	"os"
)

func main() {
	switch os.Args[1] {
	case "server":
		chat.Serve()
	case "connect":
		chat.Conn()
	default:
		log.Println("Invalid command. Use server or connect")
	}
}
