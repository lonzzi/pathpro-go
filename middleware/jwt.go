package middleware

import (
	"pathpro-go/pkg/engine"
	"pathpro-go/pkg/errno"
	"pathpro-go/utils/jwt"
	"strings"
)

func JWTAuth() engine.HandlerFunc {
	return func(c *engine.Context) *engine.Response {
		var (
			refreshToken string
			accessToken  string
			claims       *jwt.Claims
			err          error
		)

		refreshToken, err = c.Cookie("refresh_token")
		if err != nil {
			return engine.NewErrorResponse(errno.ErrRefreshTokenNotFound)
		}
		if accessToken = c.GetHeader("Authorization"); accessToken == "" {
			return engine.NewErrorResponse(errno.ErrTokenNotFound)
		}
		accessToken = strings.TrimPrefix(accessToken, "Bearer ")

		newAccessToken, newRefreshToken, err := jwt.ParseRefreshToken(accessToken, refreshToken)
		if err != nil {
			return engine.NewErrorResponse(errno.ErrTokenInvalid)
		}

		if claims, err = jwt.ParseToken(newAccessToken); err != nil {
			return engine.NewErrorResponse(errno.ErrTokenInvalid)
		}

		SetToken(c, newRefreshToken)

		c.Set("id", claims.ID)
		c.Set("username", claims.Username)

		return nil
	}
}

func SetToken(c *engine.Context, refreshToken string) {
	isSecure := IsSecure(c)
	c.SetCookie("refresh_token", refreshToken, 3600, "/", "", isSecure, true)
}

func IsSecure(c *engine.Context) bool {
	return c.GetHeader("X-Forwarded-Proto") == "https"
}
