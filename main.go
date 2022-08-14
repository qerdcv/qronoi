package main

import (
	"log"

	"github.com/qerdcv/voronoi/server"
)

func main() {
	s := server.New()

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
