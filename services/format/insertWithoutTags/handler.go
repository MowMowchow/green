package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


type Handler struct {
	sqlClient SqlClient
}

type SqlClient interface {
	Query(query string, args ...any) (*sql.Rows,error)
}

func (h *Handler) HandleInsertTransactionWithoutTag(){

	rows, err := h.sqlClient.Query("SELECT name, cad FROM transactions;")
	if err != nil {
		log.Panicln("ERROR WITH SELECT QUERY")
	}

	defer rows.Close()

	// Define a slice to hold the rows
	var rowsData [][]interface{}

	// Iterate over the rows
	for rows.Next() {
		// Create a slice to hold column values
		// var columns []interface{}
		var name, description1 string

		rows.Scan(&name, &description1)

		// Append the column values slice to the rowsData slice
		temp := []interface{}{name, description1}
		rowsData = append(rowsData, temp)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	//   Print the data
	for _, columns := range rowsData {
		fmt.Println(columns)
	}




	// transactions := make([]string, 10 )
	// scanArr := make([]interface{}, 10)
	// for i := range transactions {
	// 	scanArr[i] = &transactions[i]
	// }

	// err = rows.Scan(scanArr...)
	// if err != nil {
	// 	log.Panicln("ERROR SCANNING ROWS")
	// }

	// fmt.Println("PRINTING ROWS:")
	// for _, row := range transactions{
	// 	fmt.Println("ROW: ", row)
	// }



}

