package main

import (
	"github.com/honda-pp/SocialVueGo/backend/app"
)

func main() {
	app := app.NewApplication()
	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
