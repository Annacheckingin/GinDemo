package jwt

import (
	"GinDemo/uilty"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"strings"
	"time"
)

const JWT_SUB_CONTEXT_KEY = "JWT_SUB_CONTEXT_KEY"

func init() {
	key, err := generateECDSAKeyPair()
	if err != nil {
		panic("公私钥初始化失败")
	}
	privateKey = key
	publicKey = getPublicKey(key)
}

const theSecret = "mySecretKeyForJWT"

var publicKey *ecdsa.PublicKey
var privateKey *ecdsa.PrivateKey

func generateECDSAKeyPair() (*ecdsa.PrivateKey, error) {
	hash := sha256.Sum256([]byte(theSecret))
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.New(rand.NewSource(int64(hash[0]))))
	if err != nil {
		return nil, err
	}
	return key, nil
}

func getPublicKey(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	return &privateKey.PublicKey
}

func SimpleJwt(expire time.Duration, sub string) (string, error) {
	var (
		key *ecdsa.PrivateKey
		t   *jwt.Token
		s   string
	)
	key = privateKey
	//t = jwt.New(jwt.SigningMethodHS256)
	t = jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"exp": time.Now().Add(expire).Unix(),
			"sub": sub,
		})
	s, er := t.SignedString(key)
	if er != nil {
		return "", er
	}
	return s, nil
}

// sub不传表示不校验sub
func ValidateToken(token string, sub *string) bool {

	if sub != nil {
		jwtToken, er := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		}, jwt.WithExpirationRequired(), jwt.WithSubject(*sub))
		if er != nil {
			return false
		}
		return jwtToken.Valid
	} else {
		jwtToken, er := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		}, jwt.WithExpirationRequired())
		if er != nil {
			return false
		}
		return jwtToken.Valid
	}
}

func GetSubFromJwtToken(token string) *string {
	claims := &jwt.MapClaims{}
	tok, er := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if er != nil {
		return nil
	}
	str, err := tok.Claims.GetSubject()
	if err != nil {
		return nil
	}
	return &str
}

func SimpleJwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := c.GetHeader("Authorization")
		parts := strings.Split(payload, " ")
		if len(parts) != 2 {
			uilty.ErrorMessage(c, "token格式不对")
			c.Abort()
			return
		}
		tokenString := parts[1]
		if len(tokenString) == 0 {
			uilty.ErrorMessage(c, "token格式不对")
			c.Abort()
			return
		}
		if !ValidateToken(tokenString, nil) {
			uilty.ErrorMessage(c, "Token已过期")
			c.Abort()
			return
		}
		sub := GetSubFromJwtToken(tokenString)
		if sub == nil || len(*sub) == 0 {
			uilty.ErrorMessage(c, "令牌不正确")
			c.Abort()
			return
		}
		c.Set(JWT_SUB_CONTEXT_KEY, sub)
		c.Next()
	}
}
