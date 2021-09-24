package main

import (
	"fmt"

	"github.com/yggbrazil/go-toolbox/database"
)

func main() {
	db := database.MustGetByFile("../sqlite3-example.json")

	var now string

	err := db.Get(&now, `select current_timestamp;`)

	if err != nil {
		panic(err)
	}

	fmt.Println("current_timestamp:", now)
}
