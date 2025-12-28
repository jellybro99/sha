package sha256

import (
	"encoding/binary"
)

func preprocessMessage(messageString string) []byte {
	message := []byte(messageString)
	messageLength := uint64(len(message)) * 8

	// will need to handle cases where message isn't evenly divisable into bits
	message = append(message, 0x80)
	for len(message)*8%512 != 448 {
		message = append(message, 0x00)
	}

	var lengthBytes [8]byte
	binary.BigEndian.PutUint64(lengthBytes[:], messageLength)
	message = append(message, lengthBytes[:]...)

	return message
}

func preprocessHash() [8]uint32 {
	return [8]uint32{
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

func createBlocks(message []byte) [][16]uint32 {
	var blocks [][16]uint32

	for i := 0; i < len(message); i += 64 {
		var block [16]uint32

		for j := range 16 {
			start := i + j*4
			block[j] = binary.BigEndian.Uint32(message[start : start+4])
		}

		blocks = append(blocks, block)
	}
	return blocks
}
