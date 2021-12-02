package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	db := connectToMySQL()

	if db == nil {
		fmt.Println("Error")
	}

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello World")
	}).Methods("GET")

	log.Fatalln(http.ListenAndServe(":3576", myRouter))
}

func connectToMySQL() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/gorest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Errorf(err.Error())
	}
	return db
}
