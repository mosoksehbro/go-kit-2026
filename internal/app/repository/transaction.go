package repository

import "gorm.io/gorm"

type Transaction interface {
	WithTx(tx *gorm.DB) Transaction
}

type transaction struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) *transaction {
	return &transaction{db: db}
}

func (t *transaction) WithTx(tx *gorm.DB) Transaction {
	if tx != nil {
		return &transaction{db: tx}
	}
	return t
}
