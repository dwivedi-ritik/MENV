package menv

import (
	"crypto/rand"
	"math/big"
	"os"
)

func generateRandomString(n int) (string, error) {
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

// This method will generate new random string and store into the path config
func InitPathConfig() error {
	confPath := FetchConfigPath()
	new_key, err := generateRandomString(32)
	file, err := os.Create(confPath)
	_, err = file.WriteString(new_key)
	return err
}
