package menv

import (
	"fmt"
	"os"
)

// This method will generate new random string and store into the path config
func InitConfig() error {
	fmt.Println("Couldn't find secret key, generating new key")
	confPath := FetchConfigPath()
	new_key, err := GenerateRandomString(32) //32 byte secret key size
	file, err := os.Create(confPath)
	defer file.Close()
	_, err = file.WriteString(new_key)
	fmt.Println("New key generated")
	return err
}
