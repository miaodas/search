package hiendy

import (
	"log"
)

type Server struct {
	Status int
	Id     string
}

func (s *Server) Begin() {
	s.Status = 1
	Start()
	log.Printf("gateway server started (status: %d)", 3)
}
