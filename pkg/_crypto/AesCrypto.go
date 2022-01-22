package _crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

//AesEncrypt aes method
func AesEncrypt(data string, ConfigKey string) (string, error) {
	bKey, er := hex.DecodeString(ConfigKey)
	if er != nil {
		return "", errors.New("error")
	}
	plaintext := []byte(data)

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", errors.New("error")
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", errors.New("error")
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func AesEncryptFromByte(data []byte, ConfigKey string) (string, error) {
	bKey, er := hex.DecodeString(ConfigKey)
	if er != nil {
		return "", errors.New("error")
	}
	plaintext := data

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", errors.New("error")
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", errors.New("error")
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

//AesDecrypt aes method
func AesDecrypt(data string, ConfigKey string) (string, error) {
	ciphertext, _ := base64.URLEncoding.DecodeString(data)
	bKey, er := hex.DecodeString(ConfigKey)
	if er != nil {
		return "", errors.New("error")
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", errors.New("error")
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("error")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil

}

//AesDecryptReturnByte aes method
func AesDecryptReturnByte(data string, ConfigKey string) ([]byte, error) {
	ciphertext, _ := base64.URLEncoding.DecodeString(data)
	bKey, er := hex.DecodeString(ConfigKey)
	if er != nil {
		return nil, errors.New("error")
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return nil, errors.New("error")
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("error")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil

}
