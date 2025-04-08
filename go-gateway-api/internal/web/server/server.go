package server

import (
	"net/http"

	"github.com/Carlos6464/imersao22/go-gateway/internal/service"
	"github.com/Carlos6464/imersao22/go-gateway/internal/web/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func NewServer(accountService *service.AccountService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		port:           port,
	}
}

func (s *Server) ConfigRoute() {
	AccountHandler := handlers.NewAccountHandler(s.accountService)

	s.router.Post("/accounts", AccountHandler.Create)
	s.router.Get("/accounts", AccountHandler.Get)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
