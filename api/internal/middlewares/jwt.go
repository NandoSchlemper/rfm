package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"rfmtransportes/internal/models"
)

type JwtMiddlewareHandler func(http.ResponseWriter, *http.Request, *models.User)

type JwtMiddleware struct {
	handler JwtMiddlewareHandler
}

func SecureRoute(jwt JwtMiddlewareHandler) *JwtMiddleware {
	return &JwtMiddleware{handler: jwt}
}

func (j *JwtMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 1 Validar se temos um JWT
	// w.Header().Set("Content-Type", "application/json")
	// tokenString := r.Header.Get("Authorization")
	tokenString, err := r.Cookie("token")
	fmt.Print(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	// formattedToken := strings.TrimPrefix(tokenString, "Bearer ")
	user, err := ValidateToken(tokenString.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Sem autorização")
		return
	}

	j.handler(w, r, user)
}

func GenerateToken(payload *models.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": payload,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		},
	)
	var secret []byte = []byte(os.Getenv("SECRET_JWT"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*models.User, error) {
	var secret []byte = []byte(os.Getenv("SECRET_JWT"))
	token, err := jwt.ParseWithClaims(tokenString, &models.UserJwt{}, func(t *jwt.Token) (any, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	userClaim := token.Claims.(*models.UserJwt)
	user := userClaim.User
	return user, nil
}
