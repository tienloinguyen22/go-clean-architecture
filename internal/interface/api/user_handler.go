package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/entity"
	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/service"
	"github.com/tienloinguyen22/go-clean-architecture/pkg/httputils"
)

type UserAPIHandler struct {
	UserService service.IUserService
}

func NewUserAPIHandler(userService service.IUserService) *UserAPIHandler {
	return &UserAPIHandler{UserService: userService}
}

func (h *UserAPIHandler) HandleGetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	user, err := h.UserService.GetUserByID(ctx, id)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		httputils.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}
	if user == nil {
		httputils.ResponseWithError(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	httputils.ResonseWithJSON(w, http.StatusOK, user)
}

func (h *UserAPIHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user entity.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		httputils.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if err := h.UserService.CreateUser(ctx, &user); err != nil {
		httputils.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	httputils.ResonseWithJSON(w, http.StatusCreated, user)
}

func (h *UserAPIHandler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	var user entity.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		httputils.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	user.ID = id
	if err := h.UserService.UpdateUser(ctx, &user); err != nil {
		httputils.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	httputils.ResonseWithJSON(w, http.StatusOK, user)
}

func (h *UserAPIHandler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if err := h.UserService.DeleteUser(ctx, id); err != nil {
		httputils.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	httputils.ResonseWithJSON(w, http.StatusOK, map[string]string{"result": "user deleted successfully"})
}
