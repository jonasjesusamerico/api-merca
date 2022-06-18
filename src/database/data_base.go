package database

import (
	"api-merca/src/contexto"
	"api-merca/src/model"
	"api-merca/src/model/enum"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
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

func (d database) With() *gorm.DB {
	key := contexto.ContextoAutenticacao.GetBancoDados()
	if key == string(enum.MY_SQL) {
		return d.db_mysql
	} else if key == string(enum.POSTGRES_SQL) {
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

	StringConexaoBanco := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"password",
		"devbook",
	)
	Connection.db_mysql, err = gorm.Open(mysql.Open(StringConexaoBanco), &gorm.Config{})

	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}

	Connection.db_postgres.AutoMigrate(&model.CellPhone{}, &model.User{})
	Connection.db_mysql.AutoMigrate(&model.CellPhone{}, &model.User{})
	// DB.AutoMigrate(&models.CreditCard{})

	if err != nil {
		fmt.Println(err.Error())
	}

}
