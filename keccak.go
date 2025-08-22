package moneroutil

import (
	"math/big"

	"github.com/ebfe/keccak"
)

const (
	ChecksumLength = 4
	HashLength     = 32
)

type Hash [HashLength]byte
type Checksum [ChecksumLength]byte

func Keccak256(data ...[]byte) (result Hash) {
	h := keccak.New256()
	for _, b := range data {
		h.Write(b)
	}
	r := h.Sum(nil)
	copy(result[:], r)
	return
}

func GetChecksum(data ...[]byte) (result Checksum) {
	keccak256 := Keccak256(data...)
	copy(result[:], keccak256[:4])
	return
}

func Keccak512(data ...[]byte) (result Hash) {
	h := keccak.New512()
	for _, b := range data {
		h.Write(b)
	}
	r := h.Sum(nil)
	copy(result[:], r)
	return
}

func scReduce32(data []byte) []byte {
	l := new(big.Int)
	l.SetString("1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed", 16)
	hashInt := new(big.Int).SetBytes(data)
	reduced := new(big.Int).Mod(hashInt, l)
	result := make([]byte, 32)
	reduced.FillBytes(result) // Left-pad with zeros to 32 bytes
	return result
}
