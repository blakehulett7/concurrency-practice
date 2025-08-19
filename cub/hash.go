package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
)

var HashSecret = "9cf1f2dc6771a8beab32eefc32f50d057236d27ff3f9dcb33caf7a5aaab57132"

func GetGcm() cipher.AEAD {
	key, err := hex.DecodeString(HashSecret)
	if err != nil {
		panic("Something is wrong with the hash string")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("Something is wrong with the hash key")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic("Something is wrong with the hash key... gcm could not be created")
	}

	return gcm
}

func HashEmail(email string) string {
	gcm := GetGcm()

	nonce := make([]byte, gcm.NonceSize())
	rand.Read(nonce)

	encrypted_email := gcm.Seal(nonce, nonce, []byte(email), nil)
	hashed_email := hex.EncodeToString(encrypted_email)

	return hashed_email
}

func UnhashEmail(hash string) (string, error) {
	encrypted_data, err := hex.DecodeString(hash)
	if err != nil {
		return "", err
	}

	gcm := GetGcm()

	nonce := encrypted_data[:gcm.NonceSize()]
	encrypted_email := encrypted_data[gcm.NonceSize():]
	email_data, err := gcm.Open(nil, nonce, encrypted_email, nil)

	return string(email_data), err
}
