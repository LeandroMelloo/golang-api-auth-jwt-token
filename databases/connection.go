package databases

import (
	"github.com/leandromello/api-auth-jwt-token/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "go_auth"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

func ConnectDB() {
	dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Não foi possível connectar ao banco de dados!")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}