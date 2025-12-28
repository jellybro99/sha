package sha256

import (
	"encoding/binary"
)

func preprocessMessage(messageString string) []uint32 {
	message := []byte(messageString)
	messageLength := uint64(len(message)) * 8

	// will need to handle cases where message isn't evenly divisable into bits
	message = append(message, 0x80)
	for len(message)*8%512 != 448 {
		message = append(message, 0x00)
	}

	lengthBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(lengthBytes, messageLength)
	message = append(message, lengthBytes...)

	result := make([]uint32, len(message)/4)
	for i := range result {
		result[i] = binary.BigEndian.Uint32(message[i*4 : i*4+4])
	}
	return result
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
