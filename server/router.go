package server

import (
	"fmt"
	"matrixor/config"
	"matrixor/handlers"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	cfg    config.ServerConfig
	client *mongo.Client
	dbCfg  config.DatabaseConfig
}

func New(cfg config.ServerConfig, client *mongo.Client, dbCfg config.DatabaseConfig) *Server {
	return &Server{cfg, client, dbCfg}
}

func (s *Server) Start() {
	readingHandler := handlers.NewReadingHandler(s.client, s.dbCfg)

	mux := http.NewServeMux()
	mux.HandleFunc("/readings", readingHandler.Handle)

	port := s.cfg.Port
	if port == 0 {
		port = 8080
	}

	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      mux,
		ReadTimeout:  time.Duration(s.cfg.Timeout) * time.Second,
		WriteTimeout: time.Duration(s.cfg.Timeout) * time.Second,
	}

	fmt.Printf("%s Server listening on port %d\n", time.Now().Format(time.RFC3339), port)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
