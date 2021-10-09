package hash

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

// Compute 计算消息摘要
func Compute(alg string, message string) []byte {
	var h hash.Hash
	switch alg {
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	}

	h.Write([]byte(message))

	return h.Sum(nil)
}
