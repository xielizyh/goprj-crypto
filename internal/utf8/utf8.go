package utf8

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

func Convert(form, message string) (string, error) {
	switch form {
	case "base64":
		return base64.StdEncoding.EncodeToString([]byte(message)), nil
	case "hex":
		return hex.EncodeToString([]byte(message)), nil
		// 或者直接用格式化字符串
		// return fmt.Sprintf("%x", message), nil
	default:
		return "", errors.New("暂不支持该转换形式")
	}
}

func ConvertReverse(form, message string) (string, error) {
	var err error
	var result []byte
	switch form {
	case "base64":
		result, err = base64.StdEncoding.DecodeString(message)
	case "hex":
		result, err = hex.DecodeString(message)
	default:
		return "", errors.New("暂不支持该转换形式")
	}

	if err != nil {
		return "", err
	}
	return string(result), nil
}
