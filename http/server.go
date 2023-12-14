package httpserver

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Server is the http server used to serve requests
type Server struct {
	Address      string
	Logger       *log.Logger
	Timeout      time.Duration
	MatchService matchService
}

// Open will setup a tcp listener and serve the http requests.
func (s Server) Open() error {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		return fmt.Errorf("error opening address: %s", err.Error())
	}

	// Start HTTP server.
	server := http.Server{
		// Wrap the handler with a http.TimeoutHandler that limits the maximum
		// duration spent on a ServeHTTP call.
		// NB! The http.TimeoutHandler does not implement the http.Hijacker
		// interface and can thus not be used with WebSockets.
		Handler: http.TimeoutHandler(s.Handler(), s.Timeout, "request timeout"),
	}

	fmt.Printf("server running on localhost%s\n", s.Address)

	return fmt.Errorf("error opening address: %s", server.Serve(listener))
}

// Handler defines the HTTP layer contract.
func (s Server) Handler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	encoder := encoder{}

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		encoder.StatusResponse(ctx, w, map[string]interface{}{
			"status": "ok",
		}, http.StatusOK)
	})

	r.Route("/api", func(r chi.Router) {
		matchHandler := matchHandler{
			encoder: encoder,
			matches: s.MatchService,
		}

		r.Route("/matches", matchHandler.Routes)
	})

	return r
}
