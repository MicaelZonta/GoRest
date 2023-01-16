package main

import (
	"GoRest/config"
	"GoRest/db"
	"GoRest/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	//Carrega config
	log.Println("Start")
	err := config.Load()
	log.Println("Loaded Config")

	if err != nil {
		panic(err)
	}
	log.Println("No Error Loading Config")

	log.Println("Testing Database")
	conn, err := db.OpenConnection()
	if err != nil {
		panic(err)
	}
	conn.Close()
	log.Println("Successful Connection to Database")

	//Cria roteamento
	log.Println("Creating Routers")

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{codigo}", handlers.Update)
	r.Delete("/{codigo}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{codigo}", handlers.Get)

	log.Println("Routers Created")

	log.Println("Instanciating Router")

	//Cria Rotas
	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)
	log.Println("Instanciatiated")

}
