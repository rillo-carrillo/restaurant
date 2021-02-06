package utils

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/rillo-carrillo/restaurant/server/consts"
	"golang.org/x/crypto/bcrypt"
)

//GeneratePassword handle the Hashing of password
func GeneratePassword(pass string) (string, error) {
	bPass := []byte(pass)
	hashed, err := bcrypt.GenerateFromPassword(bPass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil

}

//CreateNewStore returns a redit Store
func CreateNewStore() redis.Store {
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(consts.GetRedisSecret()))
	if err != nil {
		fmt.Println(err)
		panic("something when wrong")
	}
	options := sessions.Options{
		MaxAge:   1000 * 60 * 60 * 24 * 7,
		HttpOnly: true,
	}
	store.Options(options)
	return store
}
