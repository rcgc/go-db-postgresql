package storage

import (
	"database/sql"
	"fmt"

	"github.com/rcgc/go-db-postgresql/pkg/invoice"
	"github.com/rcgc/go-db-postgresql/pkg/invoiceheader"
	"github.com/rcgc/go-db-postgresql/pkg/invoiceitem"
)

// MySQLInvoice used for work with MySQL - invoice
type MySQLInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems invoiceitem.Storage
}

// NewMySQLInvoice return a new pointer of MySQLInvoice
func NewMySQLInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *MySQLInvoice{
	return &MySQLInvoice{
		db: db,
		storageHeader: h,
		storageItems: i,
	}
}

// Create implement the interface invoice.Storage
func (p *MySQLInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	
	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return fmt.Errorf("header: %w", err)
	}
	fmt.Printf("Factura creada con id: %d \n", m.Header.ID)

	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return fmt.Errorf("items: %w", err)
	}
	fmt.Printf("Items creados: %d \n", len(m.Items))
	
	return tx.Commit()
}