package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	http.Server
}

func New() *Server {
	s := new(Server)
	s.Addr = ":11111"

	log.Printf("starting server at port 8080")
	http.HandleFunc("/", index)
	return s
}

func (s *Server) Run() error {
	if err := s.ListenAndServe(); err != nil {
		return fmt.Errorf("server listen: %w", err)
	}

	return nil
}
