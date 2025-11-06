package server

import (
	"blog-api/internal/database"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Run() error {
	err := s.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Close(ctx context.Context) error {
	err := s.srv.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewServer(port string, repo *database.Repo) *Server {
	r := mux.NewRouter()
	BuildHandlers(r, repo)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	return &Server{
		srv: server,
	}
}
