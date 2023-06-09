package main

import (
	"./database"
	"./promotions"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dbConnection := database.ConnectDB()

	promotions.Controller(r, dbConnection)

	err := r.Run(":8081")
	if err != nil {
		fmt.Printf("Unable to run the server: %v\n\n", err)
		return
	}
}
