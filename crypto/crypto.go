package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
)

func Encrypt(text string) string {
	key := loadKey()
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func Decrypt(cryptoText string) string {
	key := loadKey()
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext)
}

func loadKey() []byte {
	name := "bored.key"
	if f, err := ioutil.ReadFile(name); err != nil {
		k := generateKey()
		ioutil.WriteFile(name, k, 0644)
		return k
	} else {
		return f
	}
}

func generateKey() []byte {
	key := make([]byte, aes.BlockSize)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	return key
}
