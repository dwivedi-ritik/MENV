package menv

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"math/big"
	"os"
	"path/filepath"
)

func MenvExists() (bool, error) {
	current_working_dir, _ := os.Getwd()
	_, err := os.Stat(current_working_dir)
	if errors.Is(err, os.ErrNotExist) {
		return false, &FileNotExists{Err: err}
	}
	return true, nil

}

// Create parent dirs and returns the config path
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

func GetFileMetaData() {

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

	// We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
	// A nonce should always be randomly generated for every encryption.
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	// ciphertext here is actually nonce+ciphertext
	// So that when we decrypt, just knowing the nonce size
	// is enough to separate it from the ciphertext.
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return string(ciphertext)
}

func Decrypt(ciphertext string, secretKey string) string {
	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	// Since we know the ciphertext is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		panic(err)
	}

	return string(plaintext)
}
