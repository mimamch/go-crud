package handlers

import (
	"net/http"
	"strings"

	"github.com/mimamch/go-crud/internal/models"
	"github.com/mimamch/go-crud/internal/serializer"
	"github.com/mimamch/go-crud/internal/server"
)

type userHandler struct {
	Server *server.Server
}

func RegisterUserHandler(s *server.Server) {
	handler := &userHandler{
		Server: s,
	}

	s.Router.Get("/users", handler.getUserHandler)
}

func (h *userHandler) getUserHandler(w http.ResponseWriter, r *http.Request) {

	users := []models.User{}

	var values []interface{}
	var where []string

	for key, v := range r.URL.Query() {
		value := v[0]
		if value != "" {
			switch key {
			case "name":
				where = append(where, "name LIKE ?")
				values = append(values, "%"+value+"%")
			case "age":
				where = append(where, "age = ?")
				values = append(values, value)
			}
		}
	}

	if len(where) > 0 {
		h.Server.DB.Where(strings.Join(where, " AND "), values...).Find(&users)
	} else {
		h.Server.DB.Find(&users)
	}

	serializer.SendResponseData(w, users)
}
