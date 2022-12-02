package main

import (
	"planning-poker/server/db"
	"planning-poker/server/server"
)

func main() {
	db.Init()
	server.Init()
}
