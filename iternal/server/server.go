package server

import (
	"X-O-X_Game/iternal/database"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func Move(c *gin.Context) {
	room := c.Param("room")           // int -> id
	userIndex := c.Param("userIndex") //X or O
	move := c.Param("move")

	n, err := strconv.ParseInt(move, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{"mesage": "move error data"})
		log.Fatal(err)
	}

	ox := n / 10
	oy := n % 10

	db, err := database.NewPostgresDB(database.DefaultConfig)
	if err != nil {
		c.JSON(200, gin.H{"mesage": "database error data"})
		log.Fatal(err)
	}

	defer func(db *database.PostgresDB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	row, err := db.Query("SELECT place FROM gamedesc WHERE id = $1", room)
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)
	var gameplace [][]string
	for row.Next() {
		var place [][]string
		if err = row.Scan(&place); err != nil {
			c.JSON(200, gin.H{"mesage": "no room"})
			log.Fatal(err)
		}
	}
	//TODO: Create statuscode to move
	if gameplace[ox][oy] == "_" {
		_, err := db.Exec("UPDATE gamedesc SET place[$1][$2] = $3 WHERE id = $4", ox, oy, gameplace[ox][oy], room)
		if err != nil {
			c.JSON(200, gin.H{"mesage": "move error data"})
			log.Fatal(err)
		}
		gameplace[ox][oy] = userIndex
		c.JSON(200, gin.H{"mesage": "move successfully"})
	} else {
		c.JSON(200, gin.H{"mesage": "You cant do this move"})
	}

	if Checker(gameplace) == userIndex {
		c.JSON(200, gin.H{"mesage": "YOU WIN"})
	}
}
