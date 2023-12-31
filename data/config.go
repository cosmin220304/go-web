package data

import (
	"go-web/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "host=localhost user=postgres password=postgres dbname=todo_list port=5432 sslmode=disable"

func ConnectDb() error {
	var err error

	Database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&model.Todo{})

	return nil
}
