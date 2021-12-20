// Interfacer package serves the html page and handles fetching records from the
// database and displaying it to the users. It also handles user actions such as
// sending certain data back to the bluetooth module.
package main

import (
	"github.com/thinkty/warden/internal/database"
	"github.com/thinkty/warden/internal/router"
)

func main() {
	database.Init()
	router.InitAndServe()
}
