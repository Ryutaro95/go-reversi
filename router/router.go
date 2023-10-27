package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetRouting(db *sql.DB, router *gin.Engine) {
	r := router.Group("/api")
	r.POST("/games", InitStartGame(db).ServeHTTP)
}
