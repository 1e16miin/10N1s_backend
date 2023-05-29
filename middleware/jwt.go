package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CustomClaims represents the custom claims to include in the JWT token
type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// JWTAuthMiddleware is a middleware to handle JWT authentication
func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("Authorization")
		if token == "" {
			return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		jwtToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil // Replace "your-secret-key" with your own secret key
		})
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		claims, ok := jwtToken.Claims.(*CustomClaims)
		if !ok || !jwtToken.Valid {
			return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		// Store the user ID from the token in the request context
		ctx.Set("user_id", claims.UserID)

		return next(ctx)
	}
}
