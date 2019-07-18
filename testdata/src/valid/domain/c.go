package domain

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type C struct {
	ID           int            // id
	IsGlobal     int8           // is_global
	Title        sql.NullString // title
	Contents     sql.NullString // contents
	Price        sql.NullInt64  // price
	PriceNotax   sql.NullInt64  // price_notax
	DeliveryDate mysql.NullTime // delivery_date
	UpdatedAt    sql.NullInt64  // updated_at
	DeletedAt    sql.NullInt64  // deleted_at
}
