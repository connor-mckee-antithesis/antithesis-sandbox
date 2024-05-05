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

	s.router.Post("/tests/1", func(w http.ResponseWriter, r *http.Request) {
		s.count++

		// implement a race-condition sensitive operation
		s.state["steve"] += 100

		assert.Always(s.state["steve"] == 100*s.count, "state[steve] == 100 * count", nil)

		log.Printf("COMPARISON: result = %t, actual = %d, expected = %d\r\n", s.state["steve"] == 100*s.count, s.state["steve"], 100*s.count)

		w.Write([]byte(fmt.Sprintf("state[\"steve\"] == %d", s.state["steve"])))
	})

	return s
}

func (s *Server) Start() error {
	log.Println("Starting server on :8080")
	return http.ListenAndServe(":8080", s.router)
}
