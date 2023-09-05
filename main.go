package main

import (
	"flag"
	"log"

	"github.com/gofiber/template/html/v2"
	"github.com/rodmedeiross/nf-keeper/api"
)

func main() {
	listenAddr := flag.String("listen", ":8080", "server listen address")
	flag.Parse()

	engine := html.New("./templates", ".html")
	server := api.NewServer(*listenAddr, engine)

	server.InitializeRoutes()

	log.Fatal(server.Start())
}
