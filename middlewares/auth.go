package middlewares

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type jwtKeyString string

var (
	JwtCtxKey = jwtKeyString("jwt-auth-key")
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.GetHeader("Authorization")
		bearer := "Bearer "

		if auth == "" || len(auth) <= len(bearer) {
			// c.AbortWithStatus(http.StatusUnauthorized)
			// return
			c.Next()
			return
		}

		auth = auth[len(bearer):]

		validate, err := JwtValidate(auth)
		if err != nil || !validate.Valid {
			// c.AbortWithStatus(http.StatusUnauthorized)
			// return
			c.Next()
			return
		}

		claim, _ := validate.Claims.(*JwtCustomClaim)

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), JwtCtxKey, claim))

		c.Next()
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if AuthCtx(c.Request.Context()) == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func AuthCtx(ctx context.Context) *JwtCustomClaim {
	raw, _ := ctx.Value(JwtCtxKey).(*JwtCustomClaim)
	return raw
}
