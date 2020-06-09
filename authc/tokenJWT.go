package authc

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWT(experatingTime time.Duration) (string, error) {
	// Create the Claims
	StandardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * experatingTime).Unix(),
		Issuer:    "Só Por Diversão",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, StandardClaims)
	return token.SignedString([]byte("euseiquedeveriaestarusandoVariaveisAmbiente"))
}
func ValidateToken(userToken string) error {
	token, err := jwt.ParseWithClaims(userToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("euseiquedeveriaestarusandoVariaveisAmbiente"), nil
	})
	_, ok := token.Claims.(*jwt.StandardClaims)
	if ok && token.Valid {
		return nil
	}
	return err
}
