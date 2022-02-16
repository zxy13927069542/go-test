package jwt

import (
	"github.com/golang-jwt/jwt"
	log "github.com/pion/ion-log"
	"testing"
)

func TestParser(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.vSl1rHKtUsh9gCi9GXRYX27JikywMqDqY0CW6wChN9Y"
	secret := []byte("6666")
	alg := "HS256"

	//  解析token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		//  验证token签名算法
		if token.Header["alg"] != alg {
			log.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, nil
		}

		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Infof("claims: %v", claims)
	} else {
		log.Errorf("Jwt parse error! reason: %v", err)
	}
}
