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

func Proof(difficulty uint, prevHash string, txsHashString string) string {
	blockHash := ""
	for !countZero(difficulty, blockHash) {
		blockHash = doubleSHA256([]byte(prevHash + txsHashString + randString(8)))
	}
	return blockHash
}
