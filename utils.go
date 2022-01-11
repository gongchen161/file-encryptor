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

func getCipherText(fileContent []byte, hashedPassword []byte) ([]byte, error) {

	aesCipher, err := aes.NewCipher(hashedPassword)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(aesCipher)

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, fileContent, nil), nil
}

func EncryptFile(filename string, password string, outputFilename string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Unexpected error in EncryptFile", err)
		}
	}()
	fmt.Println("Start encrypting " + filename)
	fileContent, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Cannot read file " + filename)
		return
	}

	filenameByte := make([]byte, FILENAME_BLOCK_SIZE)
	copy(filenameByte, filename)
	fileContent = append(fileContent, filenameByte...)

	hashedPassword := sha256.Sum256([]byte(password))

	cipherText, err := getCipherText(fileContent, hashedPassword[:])

	if err != nil {
		fmt.Println("Encryption failed " + err.Error())
		return
	}

	err = os.WriteFile(outputFilename, cipherText, 0644)

	if err != nil {
		fmt.Println("Cannot write to output file " + outputFilename)
	} else {
		fmt.Println("Successfully encrypted " + filename + " to " + outputFilename)
	}
}

func getPlainText(cipherFileContent []byte, hashedPassword []byte) ([]byte, error) {
	aesCipher, err := aes.NewCipher(hashedPassword[:])

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(aesCipher)

	nonce, cipherText := cipherFileContent[:gcm.NonceSize()], cipherFileContent[gcm.NonceSize():]

	return gcm.Open(nil, nonce, cipherText, nil)
}

func DecryptFile(cipherFilename string, password string, outputFilename string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Unexpected error in DecryptFile", err)
		}
	}()

	fmt.Println("Start decrypting " + cipherFilename)
	cipherFileContent, err := ioutil.ReadFile(cipherFilename)

	if err != nil {
		fmt.Println("Cannot read cipher file " + cipherFilename)
		return
	}
	hashedPassword := sha256.Sum256([]byte(password))

	plainText, err := getPlainText(cipherFileContent, hashedPassword[:])

	if err != nil {
		fmt.Println("ERROR " + err.Error())
		return
	}

	plainTextSize := len(plainText) - FILENAME_BLOCK_SIZE
	if outputFilename == "" {
		outputFilename = string(plainText[plainTextSize:])
		outputFilename = strings.Replace(outputFilename, "\x00", "", -1)
	}

	err = os.WriteFile(outputFilename, plainText[:plainTextSize], 0644)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Successfully decrypted " + cipherFilename + " to " + outputFilename)
	}
}
