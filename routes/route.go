package routes

import "github.com/gin-gonic/gin"

func EnableRoute(r *gin.Engine) {
	// 应用根路由
	RootRoute(r)
	// 应用 API 路由
	ApiRoute(r)
	// 应用 Notes 路由
	NotesRoute(r)
	// 应用 Categories 路由
	CategoriesRoute(r)
	// 应用 Locations 路由
	LocationsRoute(r)
	// 应用 Search 路由
	SearchRoute(r)
	// 应用 Timeline 路由
	TimelineRoute(r)
	// 应用 Wordcloud 路由
	WordcloudRoute(r)
}
