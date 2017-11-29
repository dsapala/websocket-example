package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mccurdyc/websocket-example/server/handlers"
)

type Service struct {
	Launched time.Time
	Server   http.Server
}

func NewService(p string) *Service {
	return &Service{
		Launched: time.Now(),
		Server: http.Server{
			Addr:         p,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
	}
}

func (s *Service) Start() {
	r := mux.NewRouter()

	r.HandleFunc("/chat", handlers.Chat)
	http.Handle("/", r)

	if err := s.Server.ListenAndServe(); err != nil {
		panic(err)
	}
}
