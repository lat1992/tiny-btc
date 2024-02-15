package pow

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

func doubleSHA256(input []byte) string {
	hash1 := sha256.Sum256(input)
	hash2 := sha256.Sum256([]byte(fmt.Sprintf("%x", hash1)))
	return fmt.Sprintf("%x", hash2)
}

func countZero(n uint, hash string) bool {
	if n == 0 && len(hash) > 0 {
		return true
	}
	var count uint
	for i := 0; i < len(hash); i++ {
		if hash[i] == '0' {
			count++
		} else {
			return false
		}
		if count == n {
			return true
		}
	}
	return false
}

func randString(length uint) string {
	bytes := make([]byte, int(length))
	for i := uint(0); i < length; i++ {
		bytes[i] = byte('!' + rand.Intn('~'-'!'))
	}
	return string(bytes)
}

// Proof is the algorithm for pow, the difficulty will change according to how fast a block will be mine.
func Proof(difficulty uint, prevHash string, txsHashString string) string {
	blockHash := ""
	// count zero from the beginning of hash
	for !countZero(difficulty, blockHash) {
		// double sha254 the previous hash, txs hashes and a random string
		blockHash = doubleSHA256([]byte(prevHash + txsHashString + randString(8)))
	}
	return blockHash
}
