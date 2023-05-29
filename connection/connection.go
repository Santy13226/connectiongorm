package connection

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func GetConnection(host, user, password, dbname, port string) (*gorm.DB, error) {
	const dsn = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(fmt.Sprintf(dsn, host, user, password, dbname, port)),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true},
		})
	return db, err
}

func CloseConnection() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		err = sqlDB.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
