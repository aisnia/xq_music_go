package commons

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//返回值
type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JwtCustomClaims struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

func NewJwtCustomClaims(Id int64, username string) *JwtCustomClaims {
	return &JwtCustomClaims{
		Id: Id, UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		}}
}
func GetToken(Id int64, username string) (string, error) {
	claims := NewJwtCustomClaims(Id, username)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET))
}
