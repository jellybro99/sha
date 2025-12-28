package main

import (
	"fmt"
	"os"

	"github.com/jellybro99/sha/sha256"
)

func main() {
	message := os.Args[1]
	// check if os.Args[1] exists
	hash := sha256.Hash(message)
	fmt.Println(hash)
}
