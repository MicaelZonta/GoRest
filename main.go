package main

import (
	"GoRest/config"
	"GoRest/config/loggerconfig"
	"GoRest/config/usercontext"
	"GoRest/handlers"
	"GoRest/infra/repository/postgres/connection"
	"GoRest/port.input"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	//Setting log
	loadLogger()

	//Setting Config
	loadConfig()

	//Setting Database
	loadDatabase()

	//Setting Server
	loadServer()
}

func loadLogger() {
	loggerconfig.ConfigureBaseLogger()
}

func loadConfig() {
	//Carrega config
	log.Info().Msg("Config - START")
	err := config.Load()

	if err != nil {
		log.Err(err).Msg(fmt.Sprintf("Erro durante a busca do config, %v", err))
		panic(err)
	}
	log.Info().Msg("Config - END")
}

func loadDatabase() {
	log.Info().Msg("Database - START")
	conn, err := connection.OpenConnection()
	if err != nil {
		log.Err(err).Msg(fmt.Sprintf("Erro durante a conexão com base de dados, %v", err))
		panic(err)
	}
	conn.Close()
	log.Info().Msg("Database - END")
}

func loadServer() {
	//Cria roteamento
	log.Info().Msg("Routers - START")

	r := chi.NewRouter()
	r.Use(usercontext.CorrelationMiddleware)
	r.Post("/", port_in.CriarTarefaPortIn)
	r.Put("/{codigo}", handlers.Update)
	r.Delete("/{codigo}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{codigo}", port_in.Get)
	log.Info().Msg("Routers - END")

	//Cria Rotas
	log.Info().Msg("Server starting on port " + config.GetServerPort() + "...")
	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)
	log.Info().Msg("Server closing...")
}
