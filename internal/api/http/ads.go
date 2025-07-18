package http

import (
	"marketplace/config"
	"marketplace/internal/api/http/response"
	"marketplace/internal/api/http/types"
	"marketplace/internal/middleware"
	"marketplace/internal/usecases"
	"marketplace/pkg/http/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AdHandler struct {
	adSvc usecases.AdService
	pathCfg config.PathConfig
	svcCfg config.ServiceConfig
}

func NewAdHandler(
	adSvc usecases.AdService,
	pathCfg config.PathConfig,
	svcCfg config.ServiceConfig,
) *AdHandler {
	return &AdHandler{
		adSvc: adSvc,
		pathCfg: pathCfg,
		svcCfg: svcCfg,
	}
}

func (h *AdHandler) WithAdHandlers() handlers.RouterOption {
	return func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.WithAuthMiddleware())
			r.Post(h.pathCfg.CreateAdPath, h.postCreateAd)
		})

		r.Get(h.pathCfg.GetFeedPath, h.getFeed)
	}
}

func (h *AdHandler) postCreateAd(w http.ResponseWriter, r *http.Request) {
	req, err := types.MakePostCreateAdRequest(r, h.svcCfg)
	if err != nil {
		response.ProcessCreatingRequestError(w, err, h.svcCfg.DebugMode)
	}

	res, err := h.adSvc.CreateAd(r.Context(), &req.Ad)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
	}

	response.WriteResponse(w, map[string]string{"ad_id": res.String()}, http.StatusCreated)
}

func (h *AdHandler) getFeed(w http.ResponseWriter, r *http.Request) {
	req := types.CreateGetFeedRequest(r, h.svcCfg)

	res, err := h.adSvc.GetAdFeed(r.Context(), req.Opts)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
	}

	response.WriteResponse(w, res, http.StatusOK)
}
