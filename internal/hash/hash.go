package hash

import (
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"io"
	"os"
)

// Compute 计算消息摘要
func Compute(alg string, message string) ([]byte, error) {
	var h hash.Hash
	switch alg {
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		return nil, errors.New("暂不支持该算法")
	}

	h.Write([]byte(message))

	return h.Sum(nil), nil
}

// ComputeFile 计算文件的消息摘要
func ComputeFile(alg string, name string) ([]byte, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var h hash.Hash
	switch alg {
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		return nil, errors.New("暂不支持该算法")
	}
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
