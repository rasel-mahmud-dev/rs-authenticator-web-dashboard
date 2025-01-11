package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateBackupCodes(count int) []string {
	backupCodes := make([]string, count)
	for i := 0; i < count; i++ {
		backupCodes[i] = RandomString(12)
	}
	return backupCodes
}

func RandomString(length int) string {
	bytes := make([]byte, (length+1)/2)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("failed to generate random string: " + err.Error())
	}
	return hex.EncodeToString(bytes)[:length]
}
