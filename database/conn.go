package database

import (
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
)

// NewConn returns database conn.
func NewConn(databaseURL string) (*gorm.DB, error) {
	dialect := "mysql"
	if strings.HasSuffix(databaseURL, ".db") {
		dialect = "sqlite3"
	}

	conn, err := gorm.Open(dialect, databaseURL)
	if err != nil {
		return nil, errors.Wrap(err, "database: open connection")
	}

	return conn, nil
}
