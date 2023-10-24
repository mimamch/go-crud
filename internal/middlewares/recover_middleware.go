package middlewares

import (
	"log"
	"net/http"

	"github.com/mimamch/go-crud/internal/serializer"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				serializer.SendResponseMessage(w, 500, "Internal server error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
