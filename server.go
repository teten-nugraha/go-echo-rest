package main

import (
	"github.com/teten-nugraha/simple-go-rest/db"
	"github.com/teten-nugraha/simple-go-rest/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))

}
