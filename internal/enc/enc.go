package enc

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

// CipherEncrypt 消息加密
func CipherEncrypt(alg string, key string, message string) ([]byte, error) {
	var encKey []byte
	var err error
	if encKey, err = hex.DecodeString(key); err != nil {
		return nil, err
	}

	// 转换成hex
	plaintext, err := hex.DecodeString(message)
	if err != nil {
		return nil, err
	}

	switch alg {
	case "aes-128-ecb":
		return aesECBEncWithPKCS7Padding(encKey, plaintext)
	case "aes-128-cbc":
	case "aes-128-ctr":
	default:
		return nil, errors.New("暂不支持该算法")
	}

	return nil, nil
}

// CipherEncryptFile 加密文件
func CipherEncryptFile(alg, key, infile, outfile string) error {
	var encKey []byte
	var err error
	if encKey, err = hex.DecodeString(key); err != nil {
		return err
	}
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer out.Close()

	inReader := bufio.NewReader(in)
	outWriter := bufio.NewWriter(out)

	inbuf := make([]byte, len(encKey))
	outbuf := make([]byte, len(encKey))
	for {
		if len, err := inReader.Read(inbuf); err != nil {
			if err == io.EOF {
				break
			}
			return err
		} else {
			// 实际读取的数据
			inbuf = inbuf[:len]
		}
		switch alg {
		case "aes-128-ecb":
			outbuf, err = aesECBEncWithPKCS7Padding(encKey, inbuf)
		case "aes-128-cbc":
		case "aes-128-ctr":
		default:
			return errors.New("暂不支持该算法")
		}
		if err != nil {
			return err
		}
		if _, err := outWriter.Write(outbuf); err != nil {
			return err
		}
	}
	outWriter.Flush()

	return nil
}

// CipherEncryptFile 解密文件
func CipherDecryptFile(alg, key, infile, outfile string) error {
	var decKey []byte
	var err error
	if decKey, err = hex.DecodeString(key); err != nil {
		return err
	}
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer out.Close()

	inReader := bufio.NewReader(in)
	outWriter := bufio.NewWriter(out)

	inbuf := make([]byte, len(decKey))
	outbuf := make([]byte, len(decKey))
	for {
		if len, err := inReader.Read(inbuf); err != nil {
			if err == io.EOF {
				break
			}
			return err
		} else {
			// 实际读取的数据
			inbuf = inbuf[:len]
		}
		switch alg {
		case "aes-128-ecb":
			outbuf, err = aesECBDecWithPKCS7Padding(decKey, inbuf)
		case "aes-128-cbc":
		case "aes-128-ctr":
		default:
			return errors.New("暂不支持该算法")
		}
		if err != nil {
			return err
		}
		if _, err := outWriter.Write(outbuf); err != nil {
			return err
		}
	}
	outWriter.Flush()

	return nil
}

// CipherEncrypt 消息解密
func CipherDecrypt(alg string, key string, message string) ([]byte, error) {
	var decKey []byte
	var err error
	if decKey, err = hex.DecodeString(key); err != nil {
		return nil, err
	}

	// 转换成hex
	ciphertext, err := hex.DecodeString(message)
	if err != nil {
		return nil, err
	}

	switch alg {
	case "aes-128-ecb":
		return aesECBDecWithPKCS7Padding(decKey, ciphertext)
	case "aes-128-cbc":
	case "aes-128-ctr":
	default:
		return nil, errors.New("暂不支持该算法")
	}

	return nil, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS7UnPadding PKCS7去填充
func PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// aesECBEncWithPKCS7Padding AES-ECB模式加密，使用PKCS7填充方案
func aesECBEncWithPKCS7Padding(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
	ciphertext := make([]byte, len(plaintext))
	tmpData := make([]byte, blockSize)
	for i := 0; i < len(plaintext); i += blockSize {
		block.Encrypt(tmpData, plaintext[i:i+blockSize])
		copy(ciphertext[i:], tmpData)
	}

	return ciphertext, nil
}

// aesECBDecWithPKCS7Padding AES-ECB模式解密，使用PKCS7填充方案
func aesECBDecWithPKCS7Padding(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext := make([]byte, len(ciphertext))
	tmpData := make([]byte, blockSize)
	for i := 0; i < len(ciphertext); i += blockSize {
		block.Decrypt(tmpData, ciphertext[i:i+blockSize])
		copy(plaintext[i:], tmpData)
	}
	plaintext = PKCS7UnPadding(plaintext)

	return plaintext, nil
}
