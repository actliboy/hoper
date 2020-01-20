package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/liov/hoper/go/v2/protobuf/utils/errorcode"

	"github.com/liov/hoper/go/v2/utils/strings2"
)

type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(UserID uint64, maxAge int64, secret string) (string, error) {
	claims := Claims{
		UserID: UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + maxAge,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "hoper",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(strings2.ToBytes(secret))

	return token, err
}

func ParseToken(token, secret string) (*Claims, error) {
	tokenClaims, _ := (&jwt.Parser{SkipClaimsValidation: true}).ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return strings2.ToBytes(secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			now := time.Now().Unix()
			if claims.VerifyExpiresAt(now, false) == false {
				return nil, errorcode.LoginTimeout
			}
			return claims, nil
		}
	}

	return nil, errorcode.LoginError
}
