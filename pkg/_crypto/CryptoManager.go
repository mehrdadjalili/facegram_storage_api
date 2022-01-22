package _crypto

import (
	"crypto/md5"
	"encoding/hex"
	"facegram_file_server/config"
	"golang.org/x/crypto/sha3"
)

//AesDecryptExternalData func for decrypt data encrypted by external key.
func AesDecryptExternalData(dataEncrypt string) (string, error) {
	key := config.GetExternalKeyConfig()
	decryptData, e0 := AesDecrypt(dataEncrypt, *key)
	if e0 != nil {
		return "", e0
	}
	return decryptData, nil
}

//AesEncryptExternalData func for encrypt data by external key.
func AesEncryptExternalData(data string) (string, error) {
	key := config.GetExternalKeyConfig()
	decryptData, e0 := AesEncrypt(data, *key)
	if e0 != nil {
		return "", e0
	}
	return decryptData, nil
}

//AesDecryptInternalData func for decrypt data encrypted by internal key.
func AesDecryptInternalData(dataEncrypt string) (string, error) {
	key := config.GetInternalKeyConfig()
	decryptData, e0 := AesDecrypt(dataEncrypt, *key)
	if e0 != nil {
		return "", e0
	}
	return decryptData, nil
}

//AesEncryptInternalData func for encrypt data by internal key.
func AesEncryptInternalData(data string) (string, error) {
	key := config.GetInternalKeyConfig()
	decryptData, e0 := AesEncrypt(data, *key)
	if e0 != nil {
		return "", e0
	}
	return decryptData, nil
}

func Md5Encode(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func Sha3Encode(data string) string {
	sum := sha3.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}
