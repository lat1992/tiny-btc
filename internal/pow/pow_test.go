package pow

import (
	"crypto/sha256"
	"testing"
)

func TestCountZero(t *testing.T) {
	tests := []struct {
		name string
		n    uint
		hash string
		want bool
	}{
		{
			name: "normal",
			n:    4,
			hash: "00001234566",
			want: true,
		},
		{
			name: "more than given",
			n:    3,
			hash: "00001234566",
			want: true,
		},
		{
			name: "no zero",
			n:    0,
			hash: "1231234566",
			want: true,
		},
		{
			name: "not enough",
			n:    4,
			hash: "0001234566",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countZero(tt.n, tt.hash)
			if result != tt.want {
				t.Errorf("%s not success", tt.name)
			}
		})
	}
}

func TestProof(t *testing.T) {
	prevHash := sha256.Sum256([]byte("genesis"))
	txHash := sha256.Sum256([]byte("genesisTx"))

	t.Run("duplicate", func(t *testing.T) {
		hash1 := Proof(0, string(prevHash[0:]), string(txHash[0:]))
		hash2 := Proof(0, string(prevHash[0:]), string(txHash[0:]))
		if hash1 == hash2 {
			t.Error("duplicate hashes")
		}
	})

	t.Run("hash difficulty", func(t *testing.T) {
		hash := Proof(2, string(prevHash[0:]), string(txHash[0:]))
		if !countZero(2, hash) {
			t.Error("hash generation error")
		}
	})

	t.Run("no txs", func(t *testing.T) {
		hash := Proof(2, string(prevHash[0:]), "")
		if !countZero(2, hash) {
			t.Error("hash generation error")
		}
	})
}
