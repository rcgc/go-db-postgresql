package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/rcgc/go-db-postgresql/pkg/product"
	"github.com/rcgc/go-db-postgresql/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetByID(1); 
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("no hay producto con este id")
	case err != nil:
		log.Fatalf("product.GetByID: %v", err)
	default: 
		fmt.Println(m)
	}
	/*
	storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

	if err :=  serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceHeader.Migrate: %v", err)
	}
	*/

	/*
	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceItem.Migrate: %v", err)
	}
	*/
}