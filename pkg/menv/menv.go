package menv

import (
	"errors"
	"io"
	"os"
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
	if err != nil {
		return err
	}
	_, err = io.WriteString(menvFile, paddedString+"."+mainContent)
	return err
}

func CreateEnv(menvPath string) error {
	_, err := os.Stat(menvPath)
	if errors.Is(err, os.ErrNotExist) {
		return &FileNotExists{}
	}

	return nil
}

func performEnvCreation(menvFilePath string, secretKey string) {

}

// Takes encryption and return padded file information
func ParseFileInfo(secretKey string) {

}
