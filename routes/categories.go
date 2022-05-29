package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/routes/web"
	"timenoteWeb/utils/auth"
)

func CategoriesRoute(r *gin.Engine) {
	g := r.Group("/categories", auth.CookieTokenAuthFunc())

	g.GET("/", web.CategoryListPage)
	g.GET("/:id", web.CategoryPage)
}
