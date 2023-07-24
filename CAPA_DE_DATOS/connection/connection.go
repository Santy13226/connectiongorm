package connection

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var localDB *gorm.DB
var elephantSQLDB *gorm.DB

func GetLocalConnection() (*gorm.DB, error) {
	if localDB != nil {
		return localDB, nil
	}

	const dsn = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	var err error
	localDB, err = gorm.Open(postgres.Open(fmt.Sprintf(dsn, "localhost", "postgres", "admin123", "chatbot", "5432")),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true},
		})
	return localDB, err
}

func GetElephantSQLConnection() (*gorm.DB, error) {
	if elephantSQLDB != nil {
		return elephantSQLDB, nil
	}

	const dsn = "host=%s user=%s password=%s dbname=%s port=%s sslmode=require"
	var err error
	elephantSQLDB, err = gorm.Open(postgres.Open(fmt.Sprintf(dsn, "motty.db.elephantsql.com", "jkjilvyi", "fW_GfJ6RUn9kbYD5HmVFejQcBOyzhb96", "jkjilvyi", "5432")),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true},
		})
	return elephantSQLDB, err
}

func CloseLocalConnection() error {
	if localDB != nil {
		sqlDB, err := localDB.DB()
		if err != nil {
			return err
		}
		err = sqlDB.Close()
		if err != nil {
			return err
		}
		localDB = nil
	}
	return nil
}

func CloseElephantSQLConnection() error {
	if elephantSQLDB != nil {
		sqlDB, err := elephantSQLDB.DB()
		if err != nil {
			return err
		}
		err = sqlDB.Close()
		if err != nil {
			return err
		}
		elephantSQLDB = nil
	}
	return nil
}
