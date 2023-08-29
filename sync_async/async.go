package sync_async

import (
	"log"
	"os"
)

func Async() {
	log.Printf("Async start")
	quit := make(chan bool)

	go func() {
		PrintName()
		quit <- true
	}()

	log.Printf("Async end")
	<-quit
}

func PrintName() {
	name, err := os.ReadFile("sync_async/name.txt")
	if err != nil {
		log.Printf("Error reading sync_async/name.txt: %s", err)
		return
	}

	log.Printf("Name: %s", name)
}
