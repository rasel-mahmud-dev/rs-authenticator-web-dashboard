package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type FileSystemCache struct {
	tempDir string
}

func NewFileSystemCache() (*FileSystemCache, error) {
	dir, err := os.MkdirTemp("", "app-cache")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp dir: %v", err)
	}
	return &FileSystemCache{tempDir: dir}, nil
}

func (fs *FileSystemCache) SetItem(key string, value interface{}) error {
	filePath := filepath.Join(fs.tempDir, key+".json")
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create cache file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(value); err != nil {
		return fmt.Errorf("failed to encode value: %v", err)
	}
	return nil
}

func (fs *FileSystemCache) GetItem[T any](key string) (T, bool) {
	var zero T

	filePath := filepath.Join(fs.tempDir, key+".json")
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return zero, false
		}
		fmt.Printf("failed to open cache file: %v\n", err)
		return zero, false
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var value T
	if err := decoder.Decode(&value); err != nil {
		fmt.Printf("failed to decode cache file: %v\n", err)
		return zero, false
	}
	return value, true
}
