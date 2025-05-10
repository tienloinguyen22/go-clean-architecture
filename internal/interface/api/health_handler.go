package api

import (
	"net/http"

	"github.com/tienloinguyen22/go-clean-architecture/pkg/httputils"
)

type HeathAPIHandler struct {
	// Empty
}

func NewHeathAPIHandler() *HeathAPIHandler {
	return &HeathAPIHandler{}
}

func (h *HeathAPIHandler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	httputils.ResonseWithJSON(w, http.StatusOK, nil)
}
