package hash

import (
	"golang.org/x/crypto/bcrypt"
	"rs/auth/app/utils"
)

type HashT struct {
}

var Hash *HashT

func init() {
	Hash = &HashT{}
}

func (p *HashT) GenerateHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		return ""
	}
	return string(hash)
}

func (p *HashT) VerifyHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
