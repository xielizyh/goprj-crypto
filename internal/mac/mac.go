package mac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"os"
)

// Comute 计算消息认证码
func Compute(alg string, key string, message string) ([]byte, error) {
	var mac []byte
	var macKey []byte
	var err error
	if macKey, err = hex.DecodeString(key); err != nil {
		return nil, err
	}
	switch alg {
	case "hmac-sha256":
		m := hmac.New(sha256.New, macKey)
		m.Write([]byte(message))
		mac = m.Sum(nil)
	default:
		return nil, errors.New("暂不支持该算法")
	}

	return mac, nil
}

// ComputeFile 计算文件的消息认证码
func ComputeFile(alg string, key string, name string) ([]byte, error) {
	var macKey []byte
	var err error
	if macKey, err = hex.DecodeString(key); err != nil {
		return nil, err
	}
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var m hash.Hash
	switch alg {
	case "hmac-sha256":
		m = hmac.New(sha256.New, macKey)
	default:
		return nil, errors.New("暂不支持该算法")
	}
	if _, err := io.Copy(m, f); err != nil {
		return nil, err
	}

	return m.Sum(nil), nil
}
