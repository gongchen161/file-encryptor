package main

import (
	"crypto/sha256"
	"testing"
)

const FILE_CONTENT = "This is a plain text for testing"
const PASSWORD = "PASSWORD_TO_BE_HASHED_IN_SHA256"

func TestCorrectEncryptionDecryption(t *testing.T) {

	plainFileContentByte := ([]byte)(FILE_CONTENT)
	hashedPasswordByte := sha256.Sum256([]byte(PASSWORD))

	encryptedTextByte, err := getCipherText(plainFileContentByte, hashedPasswordByte[:])

	if err != nil {
		t.Error("Unexpected encryption error", err)
	}

	decryptedTextByte, err := getPlainText(encryptedTextByte, hashedPasswordByte[:])

	if err != nil {
		t.Error("Unexpected decryption error", err)
	}

	if (string)(decryptedTextByte) != FILE_CONTENT {
		t.Error("Incorrect results.\nExpected:", FILE_CONTENT, "\nActual:  ", (string)(decryptedTextByte))
	}

}

func TestBadHashPasswordInDecryption(t *testing.T) {

	plainFileContentByte := ([]byte)(FILE_CONTENT)
	hashedPasswordByte := sha256.Sum256([]byte(PASSWORD))

	encryptedTextByte, err := getCipherText(plainFileContentByte, hashedPasswordByte[:])

	if err != nil {
		t.Error("Unexpected encryption error", err)
	}

	// create an incorrect password in decryption
	badHashedPasswordByte := sha256.Sum256([]byte("BAD_PASSWORD"))

	_, err = getPlainText(encryptedTextByte, badHashedPasswordByte[:])

	if err == nil {
		t.Error("Error should not be nil. Wrong password used in decryption")
	}

}
