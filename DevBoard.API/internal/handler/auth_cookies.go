package handler

import (
	"net/http"
	"strings"
	"time"

	"project-devboard/internal/config"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

const (
	accessTokenCookiePath  = "/"
	refreshTokenCookiePath = "/api/v1/auth"
)

func setAuthCookies(c *gin.Context, cfg *config.Config, tokenPair *services.TokenPair) {
	if tokenPair == nil {
		return
	}

	sameSite := parseSameSite(cfg.CookieSameSite)
	accessTokenMaxAge := int(time.Until(tokenPair.ExpiresAt).Seconds())
	if accessTokenMaxAge <= 0 {
		accessTokenMaxAge = 1
	}

	writeCookie(c, cfg.AccessTokenCookieName, tokenPair.AccessToken, accessTokenMaxAge, accessTokenCookiePath, cfg.CookieDomain, cfg.CookieSecure, true, sameSite)
	writeCookie(c, cfg.RefreshTokenCookieName, tokenPair.RefreshToken, cfg.RefreshTokenExpireHours*3600, refreshTokenCookiePath, cfg.CookieDomain, cfg.CookieSecure, true, sameSite)
}

func clearAuthCookies(c *gin.Context, cfg *config.Config) {
	sameSite := parseSameSite(cfg.CookieSameSite)
	writeCookie(c, cfg.AccessTokenCookieName, "", -1, accessTokenCookiePath, cfg.CookieDomain, cfg.CookieSecure, true, sameSite)
	writeCookie(c, cfg.RefreshTokenCookieName, "", -1, refreshTokenCookiePath, cfg.CookieDomain, cfg.CookieSecure, true, sameSite)
}

func readRefreshToken(c *gin.Context, cfg *config.Config) string {
	token, err := c.Cookie(cfg.RefreshTokenCookieName)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(token)
}

func requestHasBody(c *gin.Context) bool {
	return c.Request.ContentLength > 0 || len(c.Request.TransferEncoding) > 0
}

func writeCookie(c *gin.Context, name, value string, maxAge int, path, domain string, secure, httpOnly bool, sameSite http.SameSite) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Domain:   domain,
		MaxAge:   maxAge,
		HttpOnly: httpOnly,
		Secure:   secure,
		SameSite: sameSite,
	}

	if maxAge > 0 {
		cookie.Expires = time.Now().Add(time.Duration(maxAge) * time.Second)
	} else if maxAge < 0 {
		cookie.Expires = time.Unix(0, 0)
	}

	http.SetCookie(c.Writer, cookie)
}

func parseSameSite(value string) http.SameSite {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "strict":
		return http.SameSiteStrictMode
	case "none":
		return http.SameSiteNoneMode
	case "default":
		return http.SameSiteDefaultMode
	default:
		return http.SameSiteLaxMode
	}
}

