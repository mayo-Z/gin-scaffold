package public

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func RandomUid() int {
	rand.Seed(time.Now().Unix())
	return 100000 + rand.Intn(900000)
}

func ValidPassword(afterPw, beforePw string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(afterPw), []byte(beforePw)); err != nil {
		return false
	} else {
		return true
	}
}

func SetHashedPassword(pw string) (string, error) {
	hashPw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hashPw), err
}
