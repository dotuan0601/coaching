package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	DBPostGres "github.com/tunglam268/coaching/blob/main/model.go/db"
	CoachingModel "github.com/tunglam268/coaching/blob/main/model.go/model"
)

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetTransaction")
	DB := DBPostGres.DBConnection{}
	DB = *DBPostGres.NewDBConnection()
	DB.OpenConnection()
	db := DB.ExposeDB
	var transactions []CoachingModel.Transaction
	var trans CoachingModel.Transaction
	transactions = append(transactions, trans)
	rows := db.Raw("SELECT * FROM transactions").Scan(&transactions)
	err := rows
	if err == nil {
		log.Fatal(err.Error)
	}

	fmt.Println("Transaction count: ", len(transactions))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	DB := DBPostGres.DBConnection{}
	DB = *DBPostGres.NewDBConnection()
	DB.OpenConnection()
	db := DB.ExposeDB
	var t CoachingModel.Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err == nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err1 := db.Create(&t)
	if err1 == nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err1)
	}

	json.NewEncoder(w).Encode(t)
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusCreated)

}
