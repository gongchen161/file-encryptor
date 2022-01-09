package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const FILENAME_BLOCK_SIZE = 128

func EncryptFile(filename string, password string, outputFilename string) {
	fmt.Println("Start encrypting " + filename)
	fileContent, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Cannot read file " + filename)
		return
	}

	// store and encrypt the filename
	filenameByte := make([]byte, FILENAME_BLOCK_SIZE)
	copy(filenameByte, filename)
	fileContent = append(fileContent, filenameByte...)

	passwordHash := sha256.Sum256([]byte(password))

	aesCipher, err := aes.NewCipher(passwordHash[:])

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gcm, err := cipher.NewGCM(aesCipher)

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err.Error())
		return
	}

	cipherText := gcm.Seal(nonce, nonce, fileContent, nil)

	err = os.WriteFile(outputFilename, cipherText, 0644)

	if err != nil {
		fmt.Println("Cannot write to output file " + outputFilename)
	} else {
		fmt.Println("Successfully encrypted " + filename + " to " + outputFilename)
	}
}

func DecryptFile(cipherFilename string, password string, outputFilename string) {
	fmt.Println("Start decrypting " + cipherFilename)
	fileContent, err := ioutil.ReadFile(cipherFilename)

	if err != nil {
		fmt.Println("Cannot read cipher file " + cipherFilename)
		return
	}
	passwordHash := sha256.Sum256([]byte(password))

	aesCipher, err := aes.NewCipher(passwordHash[:])

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gcm, err := cipher.NewGCM(aesCipher)

	nonce, cipherText := fileContent[:gcm.NonceSize()], fileContent[gcm.NonceSize():]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)

	if err != nil {
		fmt.Println("ERROR " + err.Error())
	}

	plainTestSize := len(plainText) - FILENAME_BLOCK_SIZE
	if outputFilename == "" {
		outputFilename = string(plainText[plainTestSize:])
		outputFilename = strings.Replace(outputFilename, "\x00", "", -1)
	}

	err = os.WriteFile(outputFilename, plainText[:plainTestSize], 0644)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Successfully decrypted " + cipherFilename + " to " + outputFilename)
	}
}
