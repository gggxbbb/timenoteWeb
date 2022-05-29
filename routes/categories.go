package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/utils/auth"
	"timenoteWeb/web"
)

func CategoriesRoute(r *gin.Engine) {
	g := r.Group("/categories", auth.CookieTokenAuthFunc())

	g.GET("/", web.CategoryListPage)
	g.GET("/:id", web.CategoryPage)
}
