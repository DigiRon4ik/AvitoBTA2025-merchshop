package jwt_token_manager

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type TokenManager struct {
	secret []byte
}

func New(secret string) *TokenManager {
	return &TokenManager{secret: []byte(secret)}
}

func (m *TokenManager) NewToken(userID, username string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		},
	})

	return token.SignedString(m.secret)
}

func (m *TokenManager) ParseToken(accessToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("incorrect signature method: %v", token.Header["alg"])
		}
		return m.secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return token, nil
}

func (m *TokenManager) ParseClaims(accessToken string) (*jwt.MapClaims, error) {
	token, err := m.ParseToken(accessToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failure when retrieving user data from token")
	}
	return &claims, nil
}
