package base64

import (
	"encoding/base64"
)

func Encode(message string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(message)), nil
}

func Decode(message string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", err
	}
	return string(decodeBytes), nil
}
