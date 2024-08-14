package menv

import (
	"errors"
	"os"
	"path/filepath"
)

const ENCRYPTION_FILE = "encryption.txt"

func GenerateMenv(env_path string) error {
	config_path := FetchConfigPath()
	encryption_file_path := filepath.Join(config_path, ENCRYPTION_FILE)
	_, err := os.Stat(encryption_file_path)
	if errors.Is(err, os.ErrNotExist) {
		return os.ErrNotExist
	}
	return nil
}
