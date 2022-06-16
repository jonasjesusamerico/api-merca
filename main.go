package main

import (
	"api-merca/src/database"
	"api-merca/src/routes"
)

func main() {
	database.ConnectWithDatabase()
	routes.Router{}.Route()
}
