package main

import (
	"X-O-X_Game/iternal/database"
	"X-O-X_Game/iternal/server"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config := database.DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "evgenii",
		Password: "password",
		DBName:   "XO",
	}

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

	router := gin.Default()
	server.SetUpRouting(router)

	router.Run(":8080")
}
