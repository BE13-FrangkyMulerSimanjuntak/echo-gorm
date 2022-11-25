package main

import (
	"frangky/be13/config"
	"frangky/be13/routes"
)

func main() {
	config.InitDB()
	config.InitialMigration()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))

}
