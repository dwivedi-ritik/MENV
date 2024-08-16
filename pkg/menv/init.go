package menv

import (
	"os"
)

// This method will generate new random string and store into the path config
func InitPathConfig() error {
	confPath := FetchConfigPath()
	new_key, err := GenerateRandomString(32) //32 byte secreat key size
	file, err := os.Create(confPath)
	defer file.Close()
	_, err = file.WriteString(new_key)
	return err
}
