package auth

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	. "timenoteWeb/config"
	. "timenoteWeb/logger"
)

type Token struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

var tokenPool []Token

func BasicAuth(c *gin.Context) bool {
	username, password, ok := c.Request.BasicAuth()
	if !ok || username != AppConfig.Admin.Username || password != AppConfig.Admin.Password {
		c.AbortWithStatus(401)
		return false
	}
	return true
}

//goland:noinspection GoUnusedExportedFunction
func BasicAuthFunc() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		AppConfig.Admin.Username: AppConfig.Admin.Password,
	})
}

func CookieTokenAuth(c *gin.Context) bool {
	Logger.Info("CookieTokenAuth: start")
	token, err := c.Cookie("token")
	if err != nil {
		Logger.Info("CookieTokenAuth: no token")
		return false
	}
	for _, t := range tokenPool {
		if t.Token == token {
			if t.ExpiresAt.Before(time.Now()) {
				Logger.Info("CookieTokenAuth: token expired")
				return false
			} else {
				Logger.Info("CookieTokenAuth: token renewed")
				t.ExpiresAt = time.Now().Add(time.Hour * 24)
				return true
			}
		}
	}
	Logger.Info("CookieTokenAuthFunc: token not found")
	return false
}

func CookieTokenAuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		Logger.Info("CookieTokenAuthFunc: start")
		token, err := c.Cookie("token")
		if err != nil {
			Logger.Info("CookieTokenAuthFunc: no token")
			c.Redirect(302, "/login")
			return
		}
		for _, t := range tokenPool {
			if t.Token == token {
				if t.ExpiresAt.Before(time.Now()) {
					Logger.Info("CookieTokenAuthFunc: token expired")
					c.Redirect(302, "/login")
					return
				} else {
					Logger.Info("CookieTokenAuthFunc: token renewed")
					t.ExpiresAt = time.Now().Add(time.Hour * 24)
					return
				}
			}
		}
		Logger.Info("CookieTokenAuthFunc: token not found")
		c.Redirect(302, "/login")
	}
}

func CheckToken(token string) bool {
	for _, t := range tokenPool {
		if t.Token == token {
			if t.ExpiresAt.Before(time.Now()) {
				Logger.Info("CheckToken: token expired")
				return false
			} else {
				Logger.Info("CheckToken: token ok")
				return true
			}
		}
	}
	return false
}

func RenewToken(token string) {
	for _, t := range tokenPool {
		if t.Token == token {
			Logger.Info("RenewToken: token renewed")
			t.ExpiresAt = time.Now().Add(time.Hour * 24)
			return
		}
	}
}

type Login struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func RequireToken(c *gin.Context) (Token, bool) {
	var data Login
	if err := c.ShouldBind(&data); err != nil {
		if err = c.ShouldBindJSON(&data); err != nil {
			if err = c.ShouldBindXML(&data); err != nil {
				Logger.Info("RequireToken: bind error")
				c.Redirect(302, "/login")
				return Token{}, false
			}
		}
	}
	if data.Username != AppConfig.Admin.Username || data.Password != AppConfig.Admin.Password {
		c.Redirect(302, "/login")
		Logger.Info("RequireToken: login failed")
		return Token{}, false
	} else {
		token := Token{
			Token:     GetRandomString(64),
			CreatedAt: time.Now(),
			ExpiresAt: time.Now().Add(time.Hour * 24),
		}
		tokenPool = append(tokenPool, token)
		c.SetCookie("token", token.Token, 180000, "/", "", false, false)
		Logger.Info("RequireToken: token created")
		return token, true
	}
}

func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
