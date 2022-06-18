package main

import (
	"api-merca/src/contexto"
	"api-merca/src/database"
	"api-merca/src/routes"
)

func init() {
	database.ConnectWithDatabase()

}

func main() {
	contexto.CriaContextoGlobalAutenticacao()
	routes.Router{}.Route()
}
