package executor

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
)

// ref: https://www.golinuxcloud.com/golang-encrypt-decrypt/#Encryption
func Decode(cryptedString string) string {
	if Prompt4K {
		Key = getSecretKey("Please enter a 32 bytes (characters) key: ")
	}
	if len(Key) != 32 {
		fmt.Printf("Current key is only %v bytes long. It needs to be of exactly 32 bytes. Aborting.\n", len(Key))
		os.Exit(1)
	}

	key := []byte(Key)
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptedString)
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
