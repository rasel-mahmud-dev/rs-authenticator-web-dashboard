package hash

import (
	"golang.org/x/crypto/bcrypt"
	"rs/auth/app/utils"
	"sync"
)

type Hash struct {
}

var (
	once sync.Once
	hash *Hash
)

var Instance = hash

func NewHash() *Hash {
	once.Do(func() {
		hash = &Hash{}
	})

	return hash
}

func init() {
	NewHash()
}

func (p *Hash) GenerateHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		return ""
	}
	return string(hash)
}

func (p *Hash) VerifyHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
