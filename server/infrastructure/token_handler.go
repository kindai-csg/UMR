package infrastructure

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type TokenHandler struct {
	secret string
}

func NewTokenHandler(secret string) *TokenHandler {
	tokenHandler := TokenHandler {
		secret: secret,
	}
	return &tokenHandler
}

func (handler *TokenHandler) CreateToken(id string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, _ := token.SignedString([]byte(handler.secret))
	return tokenString
}

func (handler *TokenHandler) AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(handler.secret), nil
	})
	
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"Msg": "認証に失敗しました",
		})
		return
	}

	claims := t.Claims.(jwt.MapClaims)
	now := time.Now().Add(time.Hour * 0).Unix()
	if (claims["exp"].(float64) < float64(now)) {
		c.AbortWithStatusJSON(400, gin.H{
			"Msg": "有効期限切れです",
		})
		return
	}
}