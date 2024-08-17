package menv

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"math/big"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

type FileMetaInfo struct {
	Name     string
	Type     string
	FileName string
}

func (f *FileMetaInfo) GetPadding() string {
	return f.Name + "," + f.Type
}

func (f *FileMetaInfo) GetFileName() string {
	if len(f.Name) > 0 && len(f.Type) > 0 {

		return f.Name + "." + f.Type
	}
	return f.Name
}

// Takes encryption and return padded file information
func ParseFileInfo(secretKey string, paddedString []byte) *FileMetaInfo {
	decryptedPadding := Decrypt(paddedString, secretKey)
	infoArr := strings.Split(decryptedPadding, ",")

	if len(infoArr) == 1 { //If env file doesn't have any extension
		return &FileMetaInfo{
			Name: infoArr[0],
		}
	}
	return &FileMetaInfo{
		Name: infoArr[0],
		Type: infoArr[1],
	}
}

func IsConfigExists() bool {
	confPath := FetchConfigPath()
	_, err := os.Stat(confPath)

	if errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil {
		panic(err)
	}

	return true
}

func IsMenvFileExists() bool {
	_, err := os.Stat(menvFile)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil {
		panic(err)
	}

	return true
}

func FetchConfigPath() string {
	homePath, _ := os.UserHomeDir()
	dirLocation := filepath.Join(homePath, ".config", "menv")
	_, err := os.Stat(dirLocation)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(dirLocation, 0755) //Octal number for file permission
	}
	return filepath.Join(dirLocation, "secretKey")
}

func FetchSecretKey() (string, error) {
	confPath := FetchConfigPath()
	content, err := os.ReadFile(confPath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func GetFileMetaData(menvFilePath string) *FileMetaInfo {
	mimeType := mime.TypeByExtension(menvFilePath)
	return &FileMetaInfo{
		Name: filepath.Base(menvFilePath),
		Type: mimeType,
	}
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func Encrypt(plaintext string, secretKey string) string {
	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext)
}

func Decrypt(ciphertext []byte, secretKey string) string {

	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), ciphertext, nil)
	if err != nil {
		panic(err)
	}
	return string(plaintext)
}
