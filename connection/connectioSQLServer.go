package connection

import (
	//"database/sql"
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetConnectionSQLS(server, database string) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("sqlserver://@%s?database=%s&trusted_connection=yes", server, database)

	db, err := gorm.Open(sqlserver.Open(connectionString),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true},
		})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings if desired
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}
