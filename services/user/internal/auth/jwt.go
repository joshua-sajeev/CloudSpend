package auth

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	UserID uint
	Role   models.UserRole
	jwt.StandardClaims
}

func GenerateToken(user *models.User) (string, error)
func ValidateToken(tokenString string) (*JWTClaims, error)
func RefreshToken(oldToken string) (string, error)
