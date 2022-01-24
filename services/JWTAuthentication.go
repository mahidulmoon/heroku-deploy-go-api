package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(email string, isUser bool, user_id int64, is_admin string) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Name   string `json:"name"`
	User   bool   `json:"user"`
	UserID int64  `json:"user_id"`
	Admin  string `json:"admin"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "mahidul",
	}
}

func getSecretKey() string {
	// secret := os.Getenv("SECRET")
	secret := "gin-insecure-f^35-g!q2%px-)_euq1_t*rc@_=(!1xwwp-j6)o6do-@-3x23z"
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(email string, isUser bool, user_id int64, is_admin string) string {
	claims := &authCustomClaims{
		email,
		isUser,
		user_id,
		is_admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
