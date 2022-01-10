package models

import "cltest/repository"

type Invoice struct {
	InvoiceId int     `db:"invoice_id"`
	InvoiceNo string  `db:"invoice_no"`
	Subtotal  float64 `db:"subtotal"`
}

func (i *Invoice) GetInvoiceById(db repository.ISQLRepository, id int) error {
	return db.GetOne(i, "SELECT invoice_id, subtotal, invoice_no FROM sales_invoice WHERE invoice_id = ?", id)
}

type InvoiceArray []*Invoice

func (i *InvoiceArray) GetLast10Invoice(db repository.ISQLRepository) error {
	getItem := func() interface{} {
		invoice := &Invoice{}
		*i = append(*i, invoice)
		return invoice
	}
	return db.GetAll(getItem, "SELECT invoice_id, subtotal, invoice_no FROM sales_invoice ORDER BY created_at DESC LIMIT 10")
}
