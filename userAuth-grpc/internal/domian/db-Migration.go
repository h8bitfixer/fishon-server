package domain

import (
	"fmt"
	"gorm.io/gorm"
)

func DatabaseTablesMigration(mysqlDB *gorm.DB) {
	mysqlDB.AutoMigrate(
		&UserAccount{},
	)

	mysqlDB.Set("gorm:table_options", "CHARSET=utf8mb4")
	mysqlDB.Set("gorm:table_options", "collation=utf8mb4_unicode_ci")

	if !mysqlDB.Migrator().HasTable(&UserAccount{}) {
		fmt.Println("CreateTable UserAccount")
		mysqlDB.Migrator().CreateTable(&UserAccount{})
	}
}
