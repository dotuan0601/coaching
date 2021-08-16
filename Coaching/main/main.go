package main

import (
	"fmt"

	DBPostGres "github.com/tunglam268/coaching/blob/main/model.go/db"
	CoachingModel "github.com/tunglam268/coaching/blob/main/model.go/model"
	"github.com/xuri/excelize/v2"
)

func main() {
	// connect to db and get connection
	DB := DBPostGres.DBConnection{}
	DB = *DBPostGres.NewDBConnection()
	DB.OpenConnection()
	db := DB.ExposeDB
	if db != nil {
		testTran := CoachingModel.Transaction{
			Id:          1,
			TraceNumber: "test",
			Block_id:    2,
			Txhash:      "test",
			Txcount:     3,
		}

		trans := []CoachingModel.Transaction{}
		trans = append(trans, testTran)
		fmt.Println(trans)

		f := excelize.NewFile()
		// Create a new sheet.
		index := f.NewSheet("Sheet1")
		// Set value of a cell.
		nametags := map[string]string{"A1": "Id", "A2": "TraceNumber", "A3": "Block_id", "A4": "Txhash", "A5": "Txcount"}
		for m, n := range nametags {
			f.SetCellValue("Sheet1", m, n)
		}
		f.SetCellValue("Sheet1", "B1", testTran.Id)
		f.SetCellValue("Sheet1", "B2", testTran.TraceNumber)
		f.SetCellValue("Sheet1", "B3", testTran.Block_id)
		f.SetCellValue("Sheet1", "B4", testTran.Txhash)
		f.SetCellValue("Sheet1", "B5", testTran.Txcount)
		// Set active sheet of the workbook.
		f.SetActiveSheet(index)
		// Save xlsx file by the given path.
		if err := f.SaveAs("Book1.xlsx"); err != nil {
			println(err.Error())
		} else {
			fmt.Println("Success to export excel")
		}

	}

}
