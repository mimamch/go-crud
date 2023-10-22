package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBUser string
	DBPass string
	DBName string
	DBHost string
	DBPort string
}

type Server struct {
	DB     *gorm.DB
	Router *chi.Mux
}

func (s *Server) Run(p string) error {
	fmt.Println("Server is running on " + p)
	return http.ListenAndServe(p, s.Router)
}

func NewServer(cfg Config) (*Server, error) {
	conn, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)))
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()

	s := &Server{
		DB:     conn,
		Router: router,
	}

	return s, nil
}
