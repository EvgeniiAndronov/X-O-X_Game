package server

import (
	"X-O-X_Game/iternal/database"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateUser(c *gin.Context) {
	userLogin := c.PostForm("userLogin")
	userName := c.PostForm("userName")
	password := c.PostForm("password")

	hashPass := makeHash(password)
	config := database.DefaultConfig
	token := makeHash(userLogin)

	type response struct {
		Id      string `json:"id"`
		Message string `json:"message"`
	}

	db, err := database.NewPostgresDB(config)
	if err != nil {
		c.JSON(http.StatusForbidden, response{Id: "-999", Message: "User has been created"})
		log.Fatal(err)
	}

	defer func(db *database.PostgresDB) {
		err := db.Close()
		if err != nil {
			c.JSON(http.StatusForbidden, response{Id: "-999", Message: "User has been created"})
			log.Fatal(err)
		}
	}(db)

	row, err := db.Query("INSERT INTO users (login, nickname, password_hash, token) VALUES ($1, $2, $3, $4) RETURNING id;", userLogin, userName, hashPass, token)
	var idUser string
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)
	for row.Next() {
		var id string
		if err := row.Scan(&id); err != nil {
			c.JSON(http.StatusForbidden, response{Id: "-999", Message: "User has been created"})
			log.Fatal(err)
		}
		idUser = id
	}

	fmt.Println(idUser)

	if err != nil {
		c.JSON(http.StatusForbidden, response{Id: "-999", Message: "User has been created"})
		log.Fatal(err)
	}

	respS := response{
		Id:      idUser,
		Message: "User suc created!",
	}

	if idUser != "" {
		c.JSON(http.StatusOK, respS)
	} else {
		respS.Message = "User fail to create!"
		c.JSON(http.StatusUnauthorized, respS)
	}

}
