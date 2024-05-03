package server

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/antithesishq/antithesis-sdk-go/assert"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	s := &Server{}

	s.router = chi.NewRouter()

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	s.router.Post("/tests/1", func(w http.ResponseWriter, r *http.Request) {
		i := rand.Intn(1e6)

		assert.Sometimes(i < 1e5, "expected i to be less than 1e5", nil)

		// adding a hard to reach assertion to demonstrate the assertion failure
		assert.Sometimes(i == 37, "expected i to be 37", nil)

		w.Write([]byte(fmt.Sprintf("%d", i)))
	})

	return s
}

func (s *Server) Start() error {
	log.Println("Starting server on :8080")
	return http.ListenAndServe(":8080", s.router)
}
