package main

import (
	"fmt"
	"database/sql"
)	
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB 
var err error

func connectToDatabase(username string, password string, database string) {
    db, err = sql.Open("mysql",  username + ":" + password + "@/" + database)
    if err != nil {
        fmt.Println(err.Error())    
    }
    // sql.DB should be long lived "defer" closes it once this function ends
    defer db.Close()

    // Test the connection to the database
    err = db.Ping()
    if err != nil {
        fmt.Println(err.Error())
    }
}