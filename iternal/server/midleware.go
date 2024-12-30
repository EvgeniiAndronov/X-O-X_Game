package server

import (
	"X-O-X_Game/iternal/database"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		validC := makeHash(takeLoginFromDBById(c.PostForm("id_user")))
		cookie, err := c.Request.Cookie("my_cookie")
		if err != nil || cookie.Value == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "cookie is missing or invalid",
			})
			return
		}

		if cookie.Value != validC {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "invalid cookie value",
			})
			return
		}
		c.Next()
	}
}

func takeLoginFromDBById(id string) string {
	config := database.DefaultConfig

	db, err := database.NewPostgresDB(config)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *database.PostgresDB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	row, err := db.Query("SELECT login FROM users WHERE id = $1 LIMIT 1;", id)
	defer row.Close()
	for row.Next() {
		var login string
		if err := row.Scan(&login); err != nil {
			log.Fatal(err)
		}
		return login
	}
	return ""
}

func CreateCookie(c *gin.Context) {
	email := c.PostForm("email")
	validCookie := makeHash(email)
	c.SetCookie("my_cookie", validCookie, 3600, "", "", false, false)

	c.JSON(http.StatusOK, gin.H{"message": "cookie set successfully"})
}

func makeHash(email string) string {
	newC := md5.Sum([]byte(email))
	return hex.EncodeToString(newC[:])
}
