package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"gopkg.in/mgo.v2/bson"
)

type user struct {
	ID   bson.ObjectId   `bson:"_id"`
	Name encryptedString `bson:"name"`
}

type encryptedString string

var key = []byte("AES256Key-32Characters1234567890")

func (s encryptedString) GetBSON() (interface{}, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(s), nil)

	fmt.Printf("To DB  : Nonce %x Ciphertext %x Full %x\n", nonce, ciphertext, append(nonce, ciphertext...))

	return append(nonce, ciphertext...), nil
}

func (s *encryptedString) SetBSON(raw bson.Raw) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	var data []byte
	if err := raw.Unmarshal(&data); err != nil {
		return err
	}

	if len(data) <= aesgcm.NonceSize() {
		return errors.New("encrypted data is too small")
	}

	nonce, ciphertext := data[:aesgcm.NonceSize()], data[aesgcm.NonceSize():]

	fmt.Printf("From DB: Nonce %x Ciphertext %x Full %x\n", nonce, ciphertext, data)

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	*s = encryptedString(string(plaintext))
	return nil
}
