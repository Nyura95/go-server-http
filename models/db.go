package models

import (
	"database/sql"
	"log"
	"messenger/config"

	// Import sql
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDb for start a pool connection
func InitDb() {
	var err error
	conf := config.GetConfig()
	db, err = sql.Open("mysql", conf.User+":"+conf.Password+"@tcp("+conf.Host+":3306)/"+conf.Database+"?parseTime=true")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

}

// Close the pool connection
func Close(db *sql.DB) {
	db.Close()
}
