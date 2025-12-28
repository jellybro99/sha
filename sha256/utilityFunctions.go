package sha256

func ch(x uint32, y uint32, z uint32) uint32 {
	return (x & y) ^ (^x & z)
}

func maj(x uint32, y uint32, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func bigSigma0(x uint32) uint32 { // Σ0
	return rotr(x, 2) ^ rotr(x, 13) ^ rotr(x, 22)
}

func bigSigma1(x uint32) uint32 { // Σ1
	return rotr(x, 6) ^ rotr(x, 11) ^ rotr(x, 25)
}

func smallSigma0(x uint32) uint32 { // σ0
	return rotr(x, 7) ^ rotr(x, 18) ^ shr(x, 3)
}

func smallSigma1(x uint32) uint32 { // σ1
	return rotr(x, 17) ^ rotr(x, 19) ^ shr(x, 10)
}

func rotr(x uint32, k uint) uint32 {
	return (x >> k) | (x << (32 - k))
}

func shr(x uint32, k uint) uint32 {
	return x >> k
}
