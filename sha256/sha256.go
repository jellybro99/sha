/*
Package sha256 implements the sha256 hash function
*/
package sha256

func Hash(messageString string) []uint32 {
	message := preprocessMessage(messageString)
	hashArray := preprocessHash()

	return hashArray
}
