package cache

import (
	"encoding/json"
	"fmt"
	"rs/auth/app/configs"
	"rs/auth/app/models"
	"rs/auth/app/utils"
)

var store = make(map[string]interface{})
var fileSystemCache *FileSystemCache
var cacheStorage string

func init() {
	cacheStorage = configs.ConfigInstance().CACHE_STORAGE
	var err error

	if cacheStorage == "filesystem" {
		fileSystemCache, err = NewFileSystemCache()
		if err != nil {
			utils.LoggerInstance.Error("Cache fs not accessible.")
		}
	}
}

func SetItem(key string, value interface{}) {

	if cacheStorage == "filesystem" {
		fileSystemCache.SetItem(key, value)
	} else {
		store[key] = value
	}
}
func GetItem[T any](key string) (T, bool) {
	var zero T
	if cacheStorage == "filesystem" {
		data, found := fileSystemCache.GetItem(key)
		if !found {
			return zero, false
		}
		var value T
		if err := json.Unmarshal(data, &value); err != nil {
			fmt.Printf("Failed to decode cache data: %v\n", err)
			return zero, false
		}
		return value, true
	} else {
		val, exists := store[key]
		if exists {
			if typedVal, ok := val.(T); ok {
				return typedVal, true
			}
		}
	}
	return zero, false
}

func GetUserFromCache(email string) *models.User {
	userC, found := GetItem[*models.User](email)
	if found {
		return userC
	}
	return nil
}
