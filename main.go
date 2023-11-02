package main

import (
	"log"

	"github.com/rcgc/go-db-postgresql/pkg/invoice"
	"github.com/rcgc/go-db-postgresql/pkg/invoiceheader"
	"github.com/rcgc/go-db-postgresql/pkg/invoiceitem"
	"github.com/rcgc/go-db-postgresql/storage"
)

func main() {
	storage.NewPostgresDB()

	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Juan",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 1},
			&invoiceitem.Model{ProductID: 2},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}
	/*
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	*/
	/*
	err := serviceProduct.Delete(3)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
	*/
	/*
	m := &product.Model{
		ID: 11,
		Name: "Curso de Go",
		Price: 50,
	}
	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}
	*/
	/*
	m, err := serviceProduct.GetByID(1); 
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("no hay producto con este id")
	case err != nil:
		log.Fatalf("product.GetByID: %v", err)
	default: 
		fmt.Println(m)
	}
	*/
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