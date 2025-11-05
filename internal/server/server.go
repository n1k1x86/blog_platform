package server

import (
	"blog-api/internal/database"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	port   string
	router *mux.Router
}

func (s *Server) Run() error {
	addr := ":" + s.port
	err := http.ListenAndServe(addr, s.router)
	if err != nil {
		return err
	}
	return nil
}

func NewServer(port string, repo *database.Repo) *Server {
	r := mux.NewRouter()
	BuildHandlers(r, repo)

	return &Server{
		port:   port,
		router: r,
	}
}
