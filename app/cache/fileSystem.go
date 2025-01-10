package cache

import (
	"encoding/json"
	"fmt"
	"io"
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

func (fs *FileSystemCache) GetItem(key string) ([]byte, bool) {
	filePath := filepath.Join(fs.tempDir, key+".json")
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false
		}
		fmt.Printf("Failed to open cache file: %v\n", err)
		return nil, false
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read cache file: %v\n", err)
		return nil, false
	}
	return data, true
}
