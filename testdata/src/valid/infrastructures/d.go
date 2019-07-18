package infrastructures

import (
	"github.com/nakamura244/dependency-check/testdata/src/valid/config"
	"github.com/nakamura244/dependency-check/testdata/src/valid/infrastructures/iface"

	// Mysql Driver
	_ "github.com/go-sql-driver/mysql"
)

// SQLhandler is struct
type SQLhandler struct {
	Conn iface.SQLAPI
}

// Result is struct
type Result struct {
	Result iface.ResultAPI
}

// Rows is struct
type Rows struct {
	Rows iface.RowsAPI
}

// Row is struct for sql.Row
type Row struct {
	Row iface.RowAPI
}

// TX is struct for tx
type TX struct {
	Tx iface.TxAPI
}

// NewSQLhandler is constructor
func NewSQLhandler(cfg *config.Config) *SQLhandler {
	return &SQLhandler{}
}
