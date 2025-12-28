/*
Package sha256 implements the sha256 hash function
*/
package sha256

func Hash(messageString string) [8]uint32 {
	message := preprocessMessage(messageString)
	hashArray := preprocessHash()

	// this is straight from nist, so I just copied their variable names
	for i := 1; i <= len(message); i++ {
		w := make([]uint32, 64)
		for t := range 64 {
			if t < 16 {
				w[t] = message[i-1]
			} else {
				w[t] = smallSigma1(w[t-2]) + w[t-7] + smallSigma0(w[t-15]) + w[t-16]
			}
		}

		var v [8]uint32
		copy(v[:], hashArray[:])

		for t := range 64 {
			t1 := v[7] + bigSigma1(v[4]) + ch(v[4], v[5], v[6]) + k[t] + w[t]
			t2 := bigSigma0(v[0]) + maj(v[0], v[1], v[2])
			v[7] = v[6]
			v[6] = v[5]
			v[5] = v[4]
			v[4] = v[3] + t1
			v[3] = v[2]
			v[2] = v[1]
			v[1] = v[0]
			v[0] = t1 + t2
		}

		for i := range hashArray {
			hashArray[i] = hashArray[i] + v[i]
		}
	}

	return hashArray
}
