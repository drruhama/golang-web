package main

import (
	"go-web-template/server"
	"go-web-template/server/configs"
	"net/http"
)

func main() {
	run()
}

func run() {
	router := http.NewServeMux()
	port := ":9999"
	db := configs.CreateConnection()
	server.StartServer(router, port, db)
}
