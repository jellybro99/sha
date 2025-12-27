package main

import (
	"encoding/binary"
)

func sha256Hash(message string) []uint32 {
	paddedMessageBytes := preprocessMessage(message)
	hashArray := preprocessHash()

	return hashArray
}

func preprocessMessage(messageString string) []byte {
	message := []byte(messageString)
	messsageLength := uint64(len(message)) * 8

	message = append(message, 0x80)
	for len(message)*8%512 != 448 {
		message = append(message, 0x00)
	}

	lengthBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(lengthBytes, messsageLength)
	message = append(message, lengthBytes...)

	return message
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
