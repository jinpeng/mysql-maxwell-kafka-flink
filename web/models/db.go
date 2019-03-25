package models

import (
	"database/sql"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitializeDatabase(connectionString string) {
	log.Println("Initializing database ...")
	var err error
	attempts := 0
	for {
		attempts++
		db, err = sql.Open("mysql", connectionString)
        if err != nil && attempts >=3 {
			log.Panic(err)
		} else if err == nil {
			break
		} else {
			log.Printf("Open attemp #%d failed, will sleep 10 seconds and try again.", attempts)
		}
		
        time.Sleep(10 * time.Second)
    }

	attempts = 0
	for {
		attempts++
		err = db.Ping()
        if err != nil && attempts >=3 {
			log.Panic(err)
		} else if err == nil {
			break
		} else {
			log.Printf("Ping attemp #%d failed, will sleep 10 seconds and try again.", attempts)
		}
		
        time.Sleep(10 * time.Second)
    }
}