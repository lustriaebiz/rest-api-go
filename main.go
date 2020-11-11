package main

import (
	"api-new/config"
	"fmt"
	"log"
	"net/http"
)

func main() {

	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")

	err := http.ListenAndServe(":2400", nil)

	if err != nil {
		log.Fatal(err)
	}

}
