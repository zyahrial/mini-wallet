package main
import (
	// "khaerus-mini-wallet/controllers"
	routes "khaerus/mini-wallet/routes"
	// "fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

func main() {
	// database.InitDB()
	// defer database.DBCon.Close()

	routes.Route()
}
