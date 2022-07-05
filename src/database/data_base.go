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

func (d database) WithContext() (bancoDados *gorm.DB) {
	key := contexto.ContextoAutenticacao.GetBancoDados()
	if key == string(enum.MY_SQL) {
		bancoDados = d.db_mysql
	} else if key == string(enum.POSTGRES_SQL) {
		bancoDados = d.db_postgres
	}
	return
}

func ConnectWithDatabase() {
	urlConexaoPostgres := "host=localhost user=admin password=admin dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	Connection.db_postgres, err = gorm.Open(postgres.Open(urlConexaoPostgres), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados 1")
	}

	StringConexaoBanco := "admin:admin@tcp(localhost:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	Connection.db_mysql, err = gorm.Open(mysql.Open(StringConexaoBanco), &gorm.Config{})

	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados Teste", err)
	}

	Connection.db_postgres.AutoMigrate(&model.Telefone{}, &model.Usuario{})
	Connection.db_mysql.AutoMigrate(&model.Telefone{}, &model.Usuario{})
	// DB.AutoMigrate(&models.CreditCard{})

	if err != nil {
		fmt.Println("Teste: ", err.Error())
	}

}
