package cache

import (
	"rs/auth/internal/models"
)

var store = make(map[string]interface{})

func SetItem(key string, value interface{}) {
	store[key] = value
}

func GetItem[T any](key string) (T, bool) {
	val, exists := store[key]
	if exists {
		if typedVal, ok := val.(T); ok {
			return typedVal, true
		}
	}
	var zero T
	return zero, false
}

func GetUserFromCache(email string) *models.User {
	userC, found := GetItem[*models.User](email)
	if found {
		return userC
	}
	return nil
}
