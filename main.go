package main

import (
	"api-new/config"
	"api-new/merchant"
	"api-new/utils"
	"context"
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

	http.HandleFunc("/merchant", GetMerchant)

	err := http.ListenAndServe(":2400", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func GetMerchant(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		merchants, err := merchant.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, merchants, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return

}
