package main

import "testing"

func TestEncryptionDecryption(t *testing.T) {
	expectedResult := "This is a plain text for testing"

	plainFileContentByte := ([]byte)(expectedResult)
	hashedPasswordByte := ([]byte)("SOME_RANDOM_HASH")

	encryptedTextByte, err := getCipherText(plainFileContentByte, hashedPasswordByte)

	if err != nil {
		t.Error("Unexpected encryption error", err)
	}

	decryptedTextByte, err := getPlainText(encryptedTextByte, hashedPasswordByte)

	if err != nil {
		t.Error("Unexpected decryption error", err)
	}

	if (string)(decryptedTextByte) == expectedResult {
		t.Error("Incorrect results.\nExpected:", expectedResult, "\nActual:  ", (string)(decryptedTextByte))
	}

}
