package driver

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func Connect() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "cheese30"
	dbName := "article"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
