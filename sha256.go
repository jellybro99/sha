package main

import (
	"encoding/binary"
	"fmt"
)

func sha256Hash(message string) []uint32 {
	paddedMessageBytes := preprocessMessage(message)
	hashArray := preprocessHash()

	fmt.Println(string(paddedMessageBytes))

	return hashArray
}

func preprocessMessage(message string) []byte {
	messageBytes := []byte(message)

	messsageLength := len(messageBytes) * 8
	paddingLength := 448 - (messsageLength+1)%512
	paddingNum := (1 << paddingLength)
	paddingBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(paddingBytes, uint64(paddingNum))

	// paddingBytes2 := make([]byte, 8)
	// binary.BigEndian.PutUint64(paddingBytes2, uint64(messsageLength))

	messageBytes = append(messageBytes, paddingBytes...)
	// messageBytes = append(messageBytes, paddingBytes2...)

	fmt.Println(message)
	fmt.Println(string(messageBytes))
	fmt.Println(messsageLength)
	fmt.Println(paddingLength)
	fmt.Println(paddingNum)
	fmt.Println(string(paddingBytes))
	fmt.Println(len(messageBytes) * 8)
	// append 1
	// append k 0s
	// append 64 bit representation of l

	paddedMessage := messageBytes

	return paddedMessage
}

func preprocessHash() []uint32 {
	return []uint32{
		0x6a09e667,
		0xbb67ae85,
		0x3c6ef372,
		0xa54ff53a,
		0x510e527f,
		0x9b05688c,
		0x1f83d9ab,
		0x5be0cd19,
	}
}
