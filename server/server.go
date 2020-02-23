package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	DB     *mongo.Client
	Router *chi.Mux
	Http   *http.Server
}

// Construct a new server
func New() *Server {
	s := &Server{}
	s.Routes()
	return s
}