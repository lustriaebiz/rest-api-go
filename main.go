package main

import (
	"api-new/config"
	"api-new/merchant"
	"api-new/models"
	"api-new/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	http.HandleFunc("/merchant/create", PostMerchant)
	http.HandleFunc("/merchant/update", UpdateMerchant)
	http.HandleFunc("/merchant/delete", DeleteMerchant)

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

func PostMerchant(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content-type application/json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		var merchants models.Merchant

		if err := json.NewDecoder(r.Body).Decode(&merchants); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}
		if err := merchant.Insert(ctx, merchants); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

func UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var merchants models.Merchant

		if err := json.NewDecoder(r.Body).Decode(&merchants); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		fmt.Println(merchants)

		if err := merchant.Update(ctx, merchants); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

func DeleteMerchant(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var merchants models.Merchant

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		merchants.ID, _ = strconv.Atoi(id)

		if err := merchant.Delete(ctx, merchants); err != nil {

			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}
