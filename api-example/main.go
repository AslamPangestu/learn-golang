package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aslampangestu/learn-golang/api-example/models"

	"github.com/aslampangestu/learn-golang/api-example/controllers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controllers.Register()
	db := models.Connect()
	// Run this in end of line
	defer db.Close()
	//Start Server
	fmt.Println("Server Start...")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
