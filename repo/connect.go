package repo

import (
	"database/sql"
	"log"
)

func connect() (db *sql.DB,err error) {
	connStr:=""
	db,err=sql.Open("postgres",connStr)
	if err!=nil{
		log.Fatal(err)
	}
	return db,err
}