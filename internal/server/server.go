package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/mimamch/go-crud/internal/middlewares"
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
	DB        *gorm.DB
	Router    *chi.Mux
	Validator *validator.Validate
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
	router.Use(middleware.Logger)
	router.Use(middlewares.RecoverMiddleware)

	s := &Server{
		DB:        conn,
		Router:    router,
		Validator: validator.New(),
	}

	return s, nil
}
