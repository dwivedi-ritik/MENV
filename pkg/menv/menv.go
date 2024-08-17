package menv

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const menvFile = "Menvfile"

type MenvFile struct {
	PaddingString   []byte
	EncryptedString []byte
}

func GetMenvFileInfo(menvFilePath string) *MenvFile {
	content, err := os.ReadFile(menvFilePath)
	if err != nil {
		panic(err)
	}
	stringContent := string(content)
	encryptedString := strings.Split(stringContent, ".")

	byteDecodedPaddedString, err := hex.DecodeString(encryptedString[0])
	if err != nil {
		panic(err)
	}

	byteDecodedEncryption, err := hex.DecodeString(encryptedString[1])
	if err != nil {
		panic(err)

	}
	return &MenvFile{
		PaddingString:   byteDecodedPaddedString,
		EncryptedString: byteDecodedEncryption,
	}

}

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
			} else if errors.Is(err, os.ErrNotExist) {
				continue
			} else {
				panic(err)
			}
		}
	} else if len(envPath) > 0 {
		_, err := os.Stat(envPath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return &FileNotExists{}
			} else {
				panic(err)
			}

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

	menvFile, err := os.Create(menvFile)
	defer menvFile.Close()
	if err != nil {
		return err
	}
	_, err = io.WriteString(menvFile, paddedString+"."+mainContent)
	return err
}

func CreateEnv(overrideFlag bool) error {
	_, err := os.Stat(menvFile)

	if errors.Is(err, os.ErrNotExist) {
		return &MenvFileNotExists{}
	}

	secretKey, err := FetchSecretKey()
	menvFileInfo := GetMenvFileInfo(menvFile)

	if overrideFlag {
		err = performEnvCreation(secretKey, menvFileInfo)
		if err != nil {
			panic(err)
		}
		return nil
	}

	fileMetaInfo := ParseFileInfo(secretKey, menvFileInfo.PaddingString)

	_, err = os.Stat(fileMetaInfo.GetFileName())

	if errors.Is(err, os.ErrNotExist) {
		err = performEnvCreation(secretKey, menvFileInfo)
		if err != nil {
			panic(err)
		}
		return nil

	}
	var overrideCommand string
	fmt.Printf("Your %v file content, will be overridden? Y/N: ", fileMetaInfo.GetFileName())
	fmt.Scanln(&overrideCommand)

	if len(overrideCommand) == 0 || strings.ToLower(overrideCommand)[0] == 'n' {
		return nil
	}

	err = performEnvCreation(secretKey, menvFileInfo)
	if err != nil {
		panic(err)
	}
	return nil
}

func performEnvCreation(secretKey string, menvFileInfo *MenvFile) error {
	fileMetaInfo := ParseFileInfo(secretKey, menvFileInfo.PaddingString)
	decryptedOriginalString := Decrypt(menvFileInfo.EncryptedString, secretKey)

	menvFile, err := os.Create(fileMetaInfo.GetFileName())
	defer menvFile.Close()
	io.WriteString(menvFile, decryptedOriginalString)
	return err
}

func UpdateMenvFile() error {
	isFileExists := IsMenvFileExists() //Not generic method, This will be checking the current executable path for Menvfile
	if !isFileExists {
		return &MenvFileNotExists{}
	}

	menvFileInfo := GetMenvFileInfo(menvFile)
	secretKey, err := FetchSecretKey()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &ConfigNotExists{}
		}
	}

	envFileInfo := ParseFileInfo(secretKey, menvFileInfo.PaddingString)
	envFileName := envFileInfo.GetFileName()
	_, err = os.Stat(envFileName)

	if errors.Is(err, os.ErrNotExist) {
		err = performMenvCreation(envFileName, secretKey)
		if err != nil {
			panic(err)
		}

	} else if err != nil {
		panic(err)
	}

	err = performMenvCreation(envFileName, secretKey)
	if err != nil {
		panic(err)
	}

	return nil
}
