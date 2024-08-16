package menv

import (
	"encoding/hex"
	"errors"
	"io"
	"os"
	"strings"
)

// Create menv file for passed env file
func CreateMenv(envPath string) error {
	possibleEnv := []string{
		".env", ".env.local",
	}
	if len(envPath) == 0 {
		for _, possibleEnv := range possibleEnv {
			_, err := os.Stat(possibleEnv)
			if err == nil {
				envPath = possibleEnv
				break
			} else if err != nil && errors.Is(err, os.ErrNotExist) {
				continue
			} else {
				panic(err)
			}
		}
	} else if len(envPath) > 0 {
		_, err := os.Stat(envPath)
		if err != nil && errors.Is(err, os.ErrNotExist) {
			return &FileNotExists{}
		} else {
			panic(err)
		}
	}

	secretKey, err := FetchSecretKey()
	if err != nil {
		panic(err)
	}
	err = performMenvCreation(envPath, secretKey)
	if err != nil {
		return err
	}
	return nil

}

func performMenvCreation(envFilePath string, secretKey string) error {
	fileInfo := GetFileMetaData(envFilePath)
	content, err := os.ReadFile(envFilePath)
	if err != nil {
		panic(err)
	}
	paddedString := Encrypt(fileInfo.GetPadding(), secretKey)

	mainContent := Encrypt(string(content), secretKey)

	menvFile, err := os.Create("menv")
	defer menvFile.Close()
	if err != nil {
		return err
	}
	_, err = io.WriteString(menvFile, paddedString+"."+mainContent)
	return err
}

func CreateEnv(menvPath string) error {
	if len(menvPath) == 0 {
		menvPath = "menv"
	}
	_, err := os.Stat(menvPath)
	if err != nil {
		panic(err)
	}
	if errors.Is(err, os.ErrNotExist) {
		return &FileNotExists{}
	}

	secretKey, err := FetchSecretKey()

	err = performEnvCreation(menvPath, secretKey)

	return err
}

func performEnvCreation(menvFilePath string, secretKey string) error {
	content, err := os.ReadFile(menvFilePath)

	if err != nil {
		return err
	}
	stringContent := string(content)

	encryptedString := strings.Split(stringContent, ".")
	paddedString := encryptedString[0]
	originalEncryptedString := encryptedString[1]

	byteDecodedEncryption, err := hex.DecodeString(originalEncryptedString)
	if err != nil {
		panic(err)
	}
	byteDecodedPaddedString, err := hex.DecodeString(paddedString)
	if err != nil {
		panic(err)

	}
	fileMetaInfo := ParseFileInfo(secretKey, byteDecodedPaddedString)
	decryptedOriginalString := Decrypt(byteDecodedEncryption, secretKey)

	menvFile, err := os.Create(fileMetaInfo.GetFileName())
	defer menvFile.Close()
	io.WriteString(menvFile, decryptedOriginalString)
	return err
}
