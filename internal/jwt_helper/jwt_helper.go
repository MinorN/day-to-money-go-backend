package jwt_helper

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateJWT(user_id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
	})
	key, err := GetHMACKey()
	if err != nil {
		log.Fatalln(err)
	}
	return token.SignedString(key)
}

func GetHMACKey() ([]byte, error) {
	keyPath := viper.GetString("jwt.hmac.key_path")
	return ioutil.ReadFile(keyPath)
}

func GenerateHMACKey() ([]byte, error) {
	key := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
