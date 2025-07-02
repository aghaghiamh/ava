package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

type MysqlDB struct {
	db *gorm.DB
}

func New(conf Config) (*MysqlDB, error) {
	connParameter := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=true", conf.Username, conf.Password, "", conf.Host, conf.Port, conf.DBName, "utf8mb4") // the empty string here define the protocol
	db, err := gorm.Open(mysql.Open(connParameter), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("database connection error: %w", err))
	}

	return &MysqlDB{
		db: db,
	}, nil
}

func (mysql *MysqlDB) GetDB() *gorm.DB {
	return mysql.db
}
