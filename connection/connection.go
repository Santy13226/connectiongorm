package connection

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	//_ "github.com/denisenkom/go-mssqldb"
)

func GetConnection(host, user, password, dbname, port string) (*gorm.DB, error) {
	const dsn = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	return gorm.Open(postgres.Open(fmt.Sprintf(dsn, host, user, password, dbname, port)),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true},
		})
}
