// Package db - содержит структуру, инициализацию и методы таблиц базы данных
package db

import (
	"database/sql"
	"github.com/azazel3ooo/yandextask/internal/models"
	"log"

	_ "github.com/lib/pq"
)

// Database - структура базы данных
type Database struct {
	Conn *sql.DB
}

// Init - инициализация базы данных и ее таблиц
func (d *Database) Init(cfg models.Config) {
	name := "postgres"

	db, err := sql.Open(name, cfg.DatabaseDsn)
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec("CREATE TABLE Users (id varchar PRIMARY KEY, urls varchar NOT NULL);")
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("CREATE TABLE Urls (id varchar PRIMARY KEY, url varchar unique NOT NULL, deleted boolean NOT NULL );")
	if err != nil {
		log.Println(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err, "from db.Init()")
	} else {
		log.Println("Connected")
	}

	d.Conn = db
}

// Ping - проверяет соединение с базой
func (d *Database) Ping() error {
	return d.Conn.Ping()
}
