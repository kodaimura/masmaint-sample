package db

import (
	//"fmt"
	"log"
	"database/sql"
	
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/lib/pq"
	//_ "github.com/go-sql-driver/mysql"

	"masmaint/config"
)


var db *sql.DB

func init() {
	var err error

	cf := config.GetConfig()

	db, err = sql.Open("sqlite3", cf.DbName)
/*
	db, err = sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cf.DbHost, cf.DbPort, cf.DbUser, cf.DbPassword, cf.DbName,
		),
	)

	db, err = sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			cf.DbUser, cf.DbPassword, cf.DbHost, cf.DbPort, cf.DbName,
		),
	)
*/
	if err != nil {
		log.Panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}