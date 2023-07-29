// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// file.go, jfgratton : 2023-07-29

package executor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

const chunkSize = 64 * 1024 // 64 KB chunk size

func encryptFile(key []byte, inputFile string, outputFile string) error {
	// Create a new AES cipher block based on the provided encryption key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Open the input file for reading
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// Create the output file for writing the encrypted data
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Generate a random IV (Initialization Vector) to use with CFB mode
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	// Write the IV to the beginning of the output file
	outFile.Write(iv)

	// Create a new CFB (Cipher Feedback) encrypter using the block cipher and IV
	stream := cipher.NewCFBEncrypter(block, iv)

	// Create a buffer to hold a chunk of data to be processed
	buf := make([]byte, chunkSize)
	for {
		// Read a chunk of data from the input file
		n, err := inFile.Read(buf)
		if n > 0 {
			// Encrypt the chunk of data in-place using XORKeyStream
			stream.XORKeyStream(buf[:n], buf[:n])

			// Write the encrypted chunk to the output file
			if _, err := outFile.Write(buf[:n]); err != nil {
				return err
			}
		}
		// Check for the end of file
		if err == io.EOF {
			break
		}
		// Handle other read errors
		if err != nil {
			return err
		}
	}

	return nil
}

func decryptFile(key []byte, inputFile string, outputFile string) error {
	// Create a new AES cipher block based on the provided encryption key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Open the input file for reading
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// Create the output file for writing the decrypted data
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Read the IV (Initialization Vector) from the beginning of the input file
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(inFile, iv); err != nil {
		return err
	}

	// Create a new CFB (Cipher Feedback) decrypter using the block cipher and IV
	stream := cipher.NewCFBDecrypter(block, iv)

	// Create a buffer to hold a chunk of data to be processed
	buf := make([]byte, chunkSize)
	for {
		// Read a chunk of data from the input file
		n, err := inFile.Read(buf)
		if n > 0 {
			// Decrypt the chunk of data in-place using XORKeyStream
			stream.XORKeyStream(buf[:n], buf[:n])

			// Write the decrypted chunk to the output file
			if _, err := outFile.Write(buf[:n]); err != nil {
				return err
			}
		}
		// Check for the end of file
		if err == io.EOF {
			break
		}
		// Handle other read errors
		if err != nil {
			return err
		}
	}

	return nil
}

func fileEncDec() {
	key := []byte("0123456789abcdef") // 16, 24, or 32 bytes for AES-128, AES-192, or AES-256 respectively

	inputFile := "input.bin"
	encryptedFile := "encrypted.bin"
	decryptedFile := "decrypted.bin"

	// Encrypt the file
	if err := encryptFile(key, inputFile, encryptedFile); err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Println("File encrypted successfully.")

	// Decrypt the file
	if err := decryptFile(key, encryptedFile, decryptedFile); err != nil {
		fmt.Println("Decryption error:", err)
		return
	}
	fmt.Println("File decrypted successfully.")
}
