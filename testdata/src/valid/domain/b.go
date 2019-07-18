package domain

import "database/sql"

type B struct {
	ID        int            // id
	Code      string         // code
	Title     sql.NullString // title
	Type      sql.NullInt64  // type
	TaxType   sql.NullInt64  // tax_type
	CreatedAt sql.NullInt64  // created_at
	UpdatedAt sql.NullInt64  // updated_at
	DeletedAt sql.NullInt64  // deleted_at
}
