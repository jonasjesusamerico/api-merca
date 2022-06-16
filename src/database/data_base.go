package database

import (
	"api-merca/src/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Connection database
	err        error
)

type database struct {
	db_mysql    *gorm.DB
	db_postgres *gorm.DB
}

func (d database) With(key string) *gorm.DB {
	if key == "MY_SQL" {
		return d.db_mysql
	} else if key == "POSTGRES" {
		return d.db_postgres
	}
	return nil
}

func ConnectWithDatabase() {
	urlConexaoPostgres := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	Connection.db_postgres, err = gorm.Open(postgres.Open(urlConexaoPostgres), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}

	// StringConexaoBanco := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
	// 	"usuario",
	// 	"senha",
	// 	"db_name",
	// )
	// Connection.db_mysql, err = gorm.Open(mysql.Open(StringConexaoBanco), &gorm.Config{})

	// if err != nil {
	// 	log.Panic("Erro ao conectar ao banco de dados")
	// }

	Connection.db_postgres.AutoMigrate(&model.CellPhone{}, &model.User{})
	// DB.AutoMigrate(&models.CreditCard{})

	if err != nil {
		fmt.Println(err.Error())
	}

}
