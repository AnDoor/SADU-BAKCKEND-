package main

import (
	"uneg.edu.ve/servicio-sadu-back/config"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()

	println("Exitted")
}
