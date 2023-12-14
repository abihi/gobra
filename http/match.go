package httpserver

import (
	"context"
	"gobra/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type matchService interface {
	Get(ctx context.Context, id uuid.UUID) (domain.Match, error)
}

type matchHandler struct {
	encoder encoder
	matches matchService
}

func (h matchHandler) Routes(router chi.Router) {
	router.Get("/", h.helloWorld)
}

func (h matchHandler) helloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp := map[string]interface{}{
		"message": "You got a match!",
	}

	h.encoder.StatusResponse(ctx, w, resp, http.StatusOK)
}
