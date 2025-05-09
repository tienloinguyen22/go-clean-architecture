package api

import (
	"net/http"

	"github.com/tienloinguyen22/go-clean-architecture/pkg/httputils"
)

type HeathHandler struct {
	// Empty
}

func NewHeathHandler() *HeathHandler {
	return &HeathHandler{}
}

func (h *HeathHandler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	httputils.ResonseWithJSON(w, http.StatusOK, nil)
}
