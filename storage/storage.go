package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func(){
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:admin@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("Connected to postgress")
	})
}

func NewMySQLDB() {
	once.Do(func(){
		var err error
		db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/godb?parseTime=true")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("Connected to mySQL")
	})
}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString{
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}