package sync_async

import (
	"log"
)

func Sync() {
	log.Printf("Sync start")
	PrintName()
	log.Printf("Sync end")
}
