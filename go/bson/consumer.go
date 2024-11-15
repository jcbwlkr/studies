package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"

	"gopkg.in/mgo.v2/bson"
)

type consumer struct {
	ID        bson.ObjectId `bson:"_id" json:"_"`
	PrivateID string        `bson:"-" json:"-"`
	PublicID  string        `bson:"public_id" json:"public_id"`
	TenantID  string        `bson:"tenant_id" json:"tenant_id"`
	Name      string        `bson:"name" json:"name"`
	NameHash  string        `bson:"name_hash" json:"-"`
}

func (c consumer) GetBSON() (interface{}, error) {
	// TODO look up key based on c.TenantID
	var key = []byte("AES256Key-32Characters1234567890")

	var err error

	// Encrypt fields and generate hashed fields. Make the hash first.
	c.NameHash = hash(c.PrivateID, c.Name)
	c.Name, err = encrypt(c.Name, key)
	if err != nil {
		return nil, err
	}

	// Return value as a new type without methods to avoid infinite recursion
	type dupe consumer
	return dupe(c), nil
}

func (c *consumer) SetBSON(raw bson.Raw) error {
	// Unmarshal the raw value into a temporary type with no SetBSON method to
	// avoid infinite recursion.
	type dupe consumer
	var tmp dupe
	if err := raw.Unmarshal(&tmp); err != nil {
		return err
	}
	// Set the unmarshaled value on the value the pointer points to. This sets
	// most of the fields.
	*c = consumer(tmp)

	// TODO look up key based on tenant ID
	var key = []byte("AES256Key-32Characters1234567890")

	// Decrypt the encrypted fields
	var err error

	c.Name, err = decrypt(c.Name, key)
	if err != nil {
		return err
	}

	// TODO get PrivateID from vault?

	return nil
}

func encrypt(s string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(s), nil)

	return string(nonce) + string(ciphertext), nil
}

func decrypt(raw string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	data := []byte(raw)

	if len(data) <= aesgcm.NonceSize() {
		// TODO or maybe it's just a normal string?
		return "", errors.New("encrypted data is too small")
	}

	nonce, ciphertext := data[:aesgcm.NonceSize()], data[aesgcm.NonceSize():]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func hash(salt, message string) string {
	data := []byte(salt)
	msg := []byte(message)
	data = append(data, msg...)
	h := sha256.Sum256(data)
	return string(h[:])
}
