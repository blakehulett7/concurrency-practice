package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
)

var HashKey = []byte("636cafd9-b647-4785-b0cd-8efd45d64f66")

func HashEmail(email string) string {
	block, err := aes.NewCipher(HashKey)
	if err != nil {
		panic("Something is wrong with the hash key")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic("Something is wrong with the hash key... gcm could not be created")
	}

	nonce := make([]byte, gcm.NonceSize())
	rand.Read(nonce)

	encrypted_email := gcm.Seal(nonce, nonce, []byte(email), nil)
	hashed_email := hex.EncodeToString(encrypted_email)

	return hashed_email
}
