package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antithesishq/antithesis-sdk-go/assert"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	state  map[string]int
	count  int
	router *chi.Mux
}

func NewServer() *Server {
	s := &Server{
		state: map[string]int{},
	}

	s.router = chi.NewRouter()

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	f := func(w http.ResponseWriter, r *http.Request) {
		s.count++

		// implement a race-condition sensitive operation
		s.state["steve"] += 100

		assert.Always(s.state["steve"] == 100*s.count, "state[steve] == 100 * count", nil)

		w.Write([]byte(fmt.Sprintf("state[\"steve\"] == %d", s.state["steve"])))
	}

	s.router.Post("/tests/1", f)

	return s
}

func (s *Server) Start() error {
	log.Println("Starting server on :8080")
	return http.ListenAndServe(":8080", s.router)
}
