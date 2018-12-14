package middleware

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DBKey string = "DB"
)

func Database(conn string) gin.HandlerFunc {
	db, err := sql.Open("sqlite3", conn)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.Set(DBKey, db)
		c.Next()
	}
}
