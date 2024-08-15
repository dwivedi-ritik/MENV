package menv

import (
	"errors"
	"os"
)

// Create menv file for passed env file
func GenerateMenv(envPath string) error {

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

	return nil

}
