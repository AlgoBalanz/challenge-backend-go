package main

import (
	"challenge/database"
	"challenge/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to SQLite database
	// db := database.GetConnection()
	// if err := database.MakeMigrations(db); err != nil {
	// 	panic(err)
	// }

	// Connect to PosgreSQL database (with GORM)
	db := database.GetConnectionGORM()
	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}

	h := handler.NewHandler(db)

	r := gin.Default()

	r.POST("/endpoint1", h.HandleEndpoint1)
	r.POST("/endpoint2", h.HandleEndpoint2)

	r.Run()
}
