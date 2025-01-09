package main

import (
	"user-management/cmd/api/modules"
)

func main() {
	app := modules.NewApp()
	app.Run()

}
