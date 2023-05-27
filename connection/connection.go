package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func GetConnection(server, user, password, database, port string) (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s",
		server, user, password, port, database)
	return sql.Open("sqlserver", connString)
}

