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
const ENCRYPT = "Encrypt"
const DECRYPT = "Decrypt"

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

func EncryptFile(filename string, password string, outputFilename string) (res string, success bool) {
	defer func() {
		if err := recover(); err != nil {
			res = "Unexpected error in EncryptFile"
			success = false
		}
	}()
	fmt.Println("Start encrypting " + filename)
	fileContent, err := ioutil.ReadFile(filename)

	if err != nil {
		res = "Cannot read file " + filename
		success = false
		return res, success
	}

	filenameByte := make([]byte, FILENAME_BLOCK_SIZE)
	copy(filenameByte, filename)
	fileContent = append(fileContent, filenameByte...)

	hashedPassword := sha256.Sum256([]byte(password))

	cipherText, err := getCipherText(fileContent, hashedPassword[:])

	if err != nil {
		res = "Encryption failed " + err.Error()
		success = false
		return res, success
	}

	err = os.WriteFile(outputFilename, cipherText, 0644)

	if err != nil {
		res = "Cannot write to output file " + outputFilename + " : " + err.Error()
		success = false
	} else {
		res = "Successfully encrypted " + filename + " to " + outputFilename
		success = true
	}
	return res, success
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

func DecryptFile(cipherFilename string, password string, outputFilename string) (res string, success bool) {
	defer func() {
		if err := recover(); err != nil {
			res = "Unexpected error in DecryptFile"
			success = false
		}
	}()

	fmt.Println("Start decrypting " + cipherFilename)
	cipherFileContent, err := ioutil.ReadFile(cipherFilename)

	if err != nil {
		res = "Cannot read cipher file " + cipherFilename
		success = false
		return res, success
	}
	hashedPassword := sha256.Sum256([]byte(password))

	plainText, err := getPlainText(cipherFileContent, hashedPassword[:])

	if err != nil {
		res = "Decryption failed " + err.Error()
		success = false
		return res, success
	}

	plainTextSize := len(plainText) - FILENAME_BLOCK_SIZE
	if outputFilename == "" {
		outputFilename = string(plainText[plainTextSize:])
		outputFilename = strings.Replace(outputFilename, "\x00", "", -1)
	}

	err = os.WriteFile(outputFilename, plainText[:plainTextSize], 0644)

	if err != nil {
		res = "Cannot write to output file " + outputFilename + " : " + err.Error()
		success = false
	} else {
		res = "Successfully decrypted " + cipherFilename + " to " + outputFilename
		success = true
	}

	return res, success
}
