package main

import (
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/db"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
