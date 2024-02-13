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

	rows, err := h.sqlClient.Query("SELECT (name) FROM transactions;")
	if err != nil {
		log.Panicln("ERROR WITH SELECT QUERY")
	}

	defer rows.Close()

	transactions := make([]string, 9 )
	scanArr := make([]interface{}, 9)
	for i := range transactions {
		scanArr[i] = &transactions[i]
	}

	err = rows.Scan(scanArr...)
	if err != nil {
		log.Panicln("ERROR SCANNING ROWS")
	}

	fmt.Println("PRINTING ROWS:")
	for _, row := range transactions{
		fmt.Println("ROW: ", row)
	}



}

