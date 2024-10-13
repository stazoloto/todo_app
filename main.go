package main

import (
	"todo_app/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
