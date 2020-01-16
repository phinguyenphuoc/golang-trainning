package main

import (
	"exercise1/db"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	Db := db.CreateDB()
	err = Db.InitDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect DB success")
	routes(Db)
}
