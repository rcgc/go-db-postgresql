package main

import (
	"github.com/rcgc/go-db-postgresql/storage"
)

func main() {
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
}