package middleware

import (
	"fmt"
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

		refreshToken = c.GetHeader("Refresh-Token")
		if accessToken = c.GetHeader("Authorization"); accessToken == "" || refreshToken == "" {
			return &engine.Response{
				Code: errno.ErrTokenNotFound,
				Msg:  errno.ErrTokenNotFound.Message,
			}
		}
		accessToken = strings.TrimPrefix(accessToken, "Bearer ")

		newAccessToken, newRefreshToken, err := jwt.ParseRefreshToken(accessToken, refreshToken)
		if err != nil {
			return &engine.Response{
				Code: errno.ErrTokenInvalid,
				Msg:  errno.ErrTokenInvalid.Message,
			}
		}

		if claims, err = jwt.ParseToken(newAccessToken); err != nil {
			return &engine.Response{
				Code: errno.ErrTokenInvalid,
				Msg:  errno.ErrTokenInvalid.Message,
			}
		}

		SetToken(c, newAccessToken, newRefreshToken)

		c.Set("id", claims.ID)
		c.Set("username", claims.Username)

		return nil
	}
}

func SetToken(c *engine.Context, accessToken, refreshToken string) {
	isSecure := IsSecure(c)
	c.Header("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	c.Header("Refresh-Token", refreshToken)
	/*
	   maxAge: cookie的有效时间，单位为秒, 0表示不保存cookie，-1表示关闭浏览器后失效
	   httpOnly: 设置为true，客户端不可读取
	*/
	c.SetCookie("Authorization", fmt.Sprintf("Bearer %s", accessToken), 3600, "/", "", isSecure, true)
	c.SetCookie("Refresh-Token", refreshToken, 3600, "/", "", isSecure, true)
}

func IsSecure(c *engine.Context) bool {
	return c.GetHeader("X-Forwarded-Proto") == "https"
}
