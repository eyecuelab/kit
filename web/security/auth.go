package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func JwtToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user": userId})
	tokenString, err := token.SignedString([]byte(viper.GetString("secret")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
