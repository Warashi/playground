package main

import (
	"log"
	"os/user"
)

func main() {
	if _, err := user.Current(); err != nil {
		log.Fatalf("user.Current: %v", err)
	}
}
