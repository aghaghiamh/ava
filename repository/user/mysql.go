package user

import (
	"github.com/aghaghiamh/ava/repository"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

// To have a consistent mysql database configs, we just use the general definition of MysqlDB struct.
func New(mysqlDB *mysql.MysqlDB) *Storage {
	db := mysqlDB.GetDB()
	db.AutoMigrate(&User{})
	return &Storage{
		db: db,
	}
}
