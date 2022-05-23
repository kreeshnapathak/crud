package middleware

import (
	"crud/services"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizatinHeaderKey   = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker services.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizatinHeader := ctx.GetHeader(authorizatinHeaderKey)
		if len(authorizatinHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		fields := strings.Fields(authorizatinHeader)
		if len(fields) < 2 {
			err := errors.New("Invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("Unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.ValidateToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
