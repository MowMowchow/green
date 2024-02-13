package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main(){

	sqlClient, err := sql.Open("mysql", "greenUser:password@tcp(ec2-34-207-157-163.compute-1.amazonaws.com:3306)/green")
	if err != nil {
		log.Println("SQL CLIENT ERR:", err)
	}

	defer sqlClient.Close()

	err = sqlClient.Ping()
	if err != nil {
		log.Println("SQL CLIENT PINF FAILED | err:", err)
	}


	handler := Handler {
		sqlClient: sqlClient,
	}

	handler.HandleInsertTransactionWithoutTag()

	
}