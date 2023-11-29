package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// *gin.Engine *gin.Contextの違いは？
func SetRouting(db *sql.DB, router *gin.Engine) {
	r := router.Group("/api")
	r.POST("/games", InitStartGame(db).ServeHTTP)
	r.GET("/games/latest/turns/:turn_count", InitLatestTurn(db).ServeHTTP)
}
