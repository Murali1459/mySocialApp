package database

import (
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type Db struct{}

func New() (*Db, error) {
	connection_string := fmt.Sprintf("%s:%s@/%s?charset=utf8", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	err := orm.RegisterDataBase("default", "mysql", connection_string)
	if err != nil {
		fmt.Errorf("Error registering db")
		return nil, err
	}
	return &Db{}, nil
}
