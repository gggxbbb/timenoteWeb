// Package auth 用于处理用户认证
//
// 提供 BasicAuth 验证 和 基于 Cookie 和 Token 的 CookieTokenAuth 验证
package auth

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	. "timenoteWeb/utils/config"
	. "timenoteWeb/utils/log"
)

// logging 包内私有 logger
var logging = Logger.WithField("包", "utils.auth")

// Token 结构体, 用于存储用户的token
type Token struct {
	// Token is the actual token
	Token string `json:"token"`
	// CreatedAt is the time the token was created
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt is the time the token expires
	ExpiresAt time.Time `json:"expires_at"`
}

// tokenPool Token 存储池
var tokenPool []Token

// BasicAuth 用于检查是否通过 BasicAuth
//
// 若通过, 将返回 true, 否则返回 false
func BasicAuth(c *gin.Context) bool {
	username, password, ok := c.Request.BasicAuth()
	if !ok || username != AppConfig.Admin.Username || password != AppConfig.Admin.Password {
		c.AbortWithStatus(401)
		return false
	}
	return true
}

// BasicAuthFunc 作为一个 gin.HandlerFunc 来验证是否通过 BasicAuth
//goland:noinspection GoUnusedExportedFunction
func BasicAuthFunc() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		AppConfig.Admin.Username: AppConfig.Admin.Password,
	})
}

// CookieTokenAuth 用于检验是否通过 CookieTokenAuth
//
// 若通过, 将返回 true, 否则返回 false
func CookieTokenAuth(c *gin.Context) bool {
	log := logging.WithField("源", "CookieTokenAuth")
	token, err := c.Cookie("token")
	if err != nil {
		log.Info("Cookies 中找不到 token")
		return false
	}
	for _, t := range tokenPool {
		if t.Token == token {
			if t.ExpiresAt.Before(time.Now()) {
				log.Info("token 已过期")
				return false
			} else {
				log.Info("token 续订")
				t.ExpiresAt = time.Now().Add(time.Hour * 24)
				return true
			}
		}
	}
	log.Info("token 不存在")
	return false
}

// CookieTokenAuthFunc 作为一个 gin.HandlerFunc 来验证是否通过 CookieTokenAuth
//goland:noinspection GoUnusedExportedFunction
func CookieTokenAuthFunc() gin.HandlerFunc {
	log := logging.WithField("源", "CookieTokenAuthFunc")
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			log.Info("Cookies 中找不到 token")
			c.Redirect(302, "/login?redirect="+c.Request.URL.Path)
		}
		for _, t := range tokenPool {
			if t.Token == token {
				if t.ExpiresAt.Before(time.Now()) {
					log.Info("token 已过期")
					c.Redirect(302, "/login?redirect="+c.Request.URL.Path)
				} else {
					log.Info("token 续订")
					t.ExpiresAt = time.Now().Add(time.Hour * 24)
				}
				return
			}
		}
		log.Info("token 不存在")
		c.Redirect(302, "/login?redirect="+c.Request.URL.Path)
	}
}

// CheckToken 用于单独检验 token 的有效性
//
// 若通过, 将返回 true, 否则返回 false
func CheckToken(token string) bool {
	log := logging.WithField("源", "CheckToken")
	for _, t := range tokenPool {
		if t.Token == token {
			if t.ExpiresAt.Before(time.Now()) {
				log.Info("token 已过期")
				return false
			} else {
				log.Info("token 正常")
				return true
			}
		}
	}
	return false
}

// RenewToken 用于延长 token 的有效期
func RenewToken(token string) {
	log := logging.WithField("源", "RenewToken")
	for _, t := range tokenPool {
		if t.Token == token {
			log.Info("token 续订成功")
			t.ExpiresAt = time.Now().Add(time.Hour * 24)
			return
		}
	}
}

// Login 用于存储登录表单的数据
type Login struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// RequireToken 用于在登陆时获取 token
//
// 若成功登录, 将返回 true, 否则返回 false
func RequireToken(c *gin.Context) bool {
	log := logging.WithField("源", "RequireToken")
	var data Login
	if err := c.ShouldBind(&data); err != nil {
		log.Info("登陆数据解析失败")
		c.Redirect(302, "/login?redirect="+c.Request.URL.Path)
		return false
	}
	if data.Username != AppConfig.Admin.Username || data.Password != AppConfig.Admin.Password {
		c.Redirect(302, "/login?redirect="+c.Request.URL.Path)
		log.Info("登陆失败")
		return false
	} else {
		token := Token{
			Token:     GetRandomString(64),
			CreatedAt: time.Now(),
			ExpiresAt: time.Now().Add(time.Hour * 24),
		}
		tokenPool = append(tokenPool, token)
		c.SetCookie("token", token.Token, 180000, "/", "", false, false)
		log.Info("token 已创建")
		return true
	}
}

// GetRandomString 用于获取随机字符串作为 token 使用
func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
