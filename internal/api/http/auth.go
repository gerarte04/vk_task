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

// @Summary 	Register new user
// @Description REGISTER NEW USER
// @Tags 		auth
// @Accept  	json
// @Produce 	json
// @Param 		credentials body 	types.PostRegisterRequest true "Login and password"
// @Success 	201 {object} 		types.PostRegisterResponse "Successfully registered"
// @Failure 	400 {string} 		string "Bad request"
// @Failure 	401 {string} 		string "Unauthorized"
// @Failure 	404 {string} 		string "Object not found"
// @Failure 	500 {string} 		string "Internal error"
// @Router		/v1/auth/register 	[post]
func (h *AuthHandler) postRegister(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostRegisterRequest(r, h.svcCfg)
	if err != nil {
		response.ProcessCreatingRequestError(w, err, h.svcCfg.DebugMode)
		return
	}

	res, err := h.authSvc.Register(r.Context(), &domain.User{Login: req.Login}, req.Password)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
		return
	}

	response.WriteResponse(w, types.PostRegisterResponse{UserId: res.String()}, http.StatusCreated)
}

// @Summary 	Login and get access token
// @Tags 		auth
// @Accept  	json
// @Produce 	json
// @Param 		credentials body 	types.PostLoginRequest true "Login and password"
// @Success 	200 {object} 		types.PostLoginResponse "Successfully authorized"
// @Failure 	400 {string} 		string "Bad request"
// @Failure 	401 {string} 		string "Unauthorized"
// @Failure 	404 {string} 		string "Object not found"
// @Failure 	500 {string} 		string "Internal error"
// @Router		/v1/auth/login 		[post]
func (h *AuthHandler) postLogin(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostLoginRequest(r)
	if err != nil {
		response.ProcessCreatingRequestError(w, err, h.svcCfg.DebugMode)
		return
	}

	res, err := h.authSvc.Login(r.Context(), req.Login, req.Password)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
		return
	}

	response.WriteResponse(w, types.PostLoginResponse{Token: res}, http.StatusOK)
}
