package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	host := "127.0.0.1"
	username := "hppoc"
	password := "password"
	dbname := "fabricexplorer"
	port := "5432"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connect Success !")
	}

	for index := range trans {
		db.Create(&trans[index])
	}
	var transactions *Transactions
	id := 1
	db.Raw("SELECT * FROM transactions WHERE id = ?", id).Scan(transactions)
}
