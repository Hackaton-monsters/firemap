package main

import (
	"firemap/cmd/command"
	"log"
)

func main() {
	if err := command.Execute(); err != nil {
		log.Fatalf("Terminate: %s", err)
	}
}
