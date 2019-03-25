package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"hello-element-vue-backend/config"
	"log"
)

var db *sqlx.DB

func InitDatabase() {
	var err error
	db, err = sqlx.Connect("sqlite3", config.Conf.SQLiteFile)
	if err != nil {
		log.Fatal(err)
	}
}
