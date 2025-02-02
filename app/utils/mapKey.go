package utils

func MapKey(data map[string]interface{}, keys ...string) interface{} {
	var current interface{} = data
	for _, key := range keys {
		if m, ok := current.(map[string]interface{}); ok {
			current = m[key]
		} else {
			return nil
		}
	}
	return current
}
