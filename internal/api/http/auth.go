package http

import (
	"marketplace/config"
	"marketplace/internal/api/http/response"
	"marketplace/internal/api/http/types"
	"marketplace/internal/domain"
	"marketplace/internal/usecases"
	"marketplace/pkg/http/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	authSvc usecases.AuthService
	pathCfg config.PathConfig
	svcCfg config.ServiceConfig
}

func NewAuthHandler(
	authSvc usecases.AuthService,
	pathCfg config.PathConfig,
	svcCfg config.ServiceConfig,
) *AuthHandler {
	return &AuthHandler{
		authSvc: authSvc,
		pathCfg: pathCfg,
		svcCfg: svcCfg,
	}
}

func (h *AuthHandler) WithAuthHandlers() handlers.RouterOption {
	return func(r chi.Router) {
		r.Post(h.pathCfg.RegisterPath, h.postRegister)
		r.Post(h.pathCfg.LoginPath, h.postLogin)
	}
}

func (h *AuthHandler) postRegister(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostRegisterRequest(r, h.svcCfg)
	if err != nil {
		response.ProcessCreatingRequestError(w, err, h.svcCfg.DebugMode)
	}

	res, err := h.authSvc.Register(r.Context(), &domain.User{Login: req.Login}, req.Password)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
	}

	response.WriteResponse(w, map[string]string{"user_id": res.String()}, http.StatusCreated)
}

func (h *AuthHandler) postLogin(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostLoginRequest(r)
	if err != nil {
		response.ProcessCreatingRequestError(w, err, h.svcCfg.DebugMode)
	}

	res, err := h.authSvc.Login(r.Context(), req.Login, req.Password)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
	}

	response.WriteResponse(w, map[string]string{"token": res}, http.StatusOK)
}
