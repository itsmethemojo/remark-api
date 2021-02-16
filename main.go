package main

import (
	bookmarkRepository "./repositories/bookmark"
	"./routes"
)

func main() {
	bookmarkRepository.InitializeDatabase()
	routes.Run()
}
