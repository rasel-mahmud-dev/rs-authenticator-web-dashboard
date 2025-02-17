package cache

import (
	"encoding/json"
	"fmt"
	"rs/auth/app/configs"
	"rs/auth/app/models"
	"rs/auth/app/utils"
	"time"
)

type Cache[T any] struct {
	CreatedAt time.Time `json:"created_at"`
	Data      *T        `json:"data"`
}

var store = make(map[string]Cache[interface{}])
var fileSystemCache *FileSystemCache
var cacheStorage string

func init() {
	cacheStorage = configs.Config.CACHE_STORAGE
	var err error

	if cacheStorage == "filesystem" {
		fileSystemCache, err = NewFileSystemCache()
		if err != nil {
			utils.LoggerInstance.Error("Cache fs not accessible.")
		}
	}
}

func SetItem(key string, value interface{}) {
	go func(key string, value interface{}) {
		newEntity := Cache[interface{}]{
			CreatedAt: time.Now(),
			Data:      &value,
		}
		if cacheStorage == "filesystem" {
			fileSystemCache.SetItem(key, newEntity)
		} else {
			store[key] = newEntity
		}
	}(key, value)
}

func GetItem[T any](key string) Cache[T] {
	var zero Cache[T]

	if cacheStorage == "filesystem" {
		data, found := fileSystemCache.GetItem(key)
		if !found {
			return zero
		}
		var value Cache[T]
		if err := json.Unmarshal(data, &value); err != nil {
			fmt.Printf("Failed to decode cache data: %v\n", err)
			return zero
		}

		return Cache[T]{
			CreatedAt: value.CreatedAt,
			Data:      value.Data,
		}
	} else {
		// ... other caching storages
	}
	return zero
}

func GetUserFromCache(email string) *models.User {
	userC := GetItem[*models.User](email)
	if userC.Data != nil {
		return *userC.Data
	}
	return nil
}
