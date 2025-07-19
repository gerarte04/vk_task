package http

import (
	"fmt"
	"marketplace/config"
	"marketplace/internal/api/http/response"
	"marketplace/internal/api/http/types"
	"marketplace/internal/middleware"
	"marketplace/internal/usecases"
	"marketplace/pkg/http/handlers"
	pkgCrypto "marketplace/pkg/utils/crypto"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AdHandler struct {
	adSvc usecases.AdService
	pathCfg config.PathConfig
	svcCfg config.ServiceConfig
	publicKey []byte
}

func NewAdHandler(
	adSvc usecases.AdService,
	pathCfg config.PathConfig,
	svcCfg config.ServiceConfig,
	publicKeyPEM string,
) (*AdHandler, error) {
	const op = "NewAdHandler"

	publicKey, err := pkgCrypto.ParsePublicKeyFromPEM(publicKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &AdHandler{
		adSvc: adSvc,
		pathCfg: pathCfg,
		svcCfg: svcCfg,
		publicKey: publicKey,
	}, nil
}

func (h *AdHandler) WithAdHandlers() handlers.RouterOption {
	return func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.WithAuthMiddleware(h.publicKey, h.svcCfg.DebugMode))
			r.Post(h.pathCfg.CreateAdPath, h.postCreateAd)
		})

		r.Get(h.pathCfg.GetFeedPath, h.getFeed)
	}
}

// @Summary 	Create new ad
// @Tags 		ads
// @Accept  	json
// @Produce 	json
// @Param		Authorization 	header	string true "Acess token with Bearer prefix"
// @Param 		ad 	body 				domain.Ad true "Ad details"
// @Success 	201 {object} 			types.PostCreateAdResponse "Successfully created"
// @Failure 	400 {string} 			string "Bad request"
// @Failure 	401 {string} 			string "Unauthorized"
// @Failure 	404 {string} 			string "Object not found"
// @Failure 	500 {string} 			string "Internal error"
// @Router		/v1/ads/create 	[post]
func (h *AdHandler) postCreateAd(w http.ResponseWriter, r *http.Request) {
	req, err := types.MakePostCreateAdRequest(r, h.svcCfg)
	if err != nil {
		response.ProcessCreatingRequestError(w, err, h.svcCfg.DebugMode)
		return
	}

	res, err := h.adSvc.CreateAd(r.Context(), &req.Ad)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
		return
	}

	response.WriteResponse(w, types.PostCreateAdResponse{AdId: res.String()}, http.StatusCreated)
}

// @Summary 	Get feed with options
// @Tags 		ads
// @Accept  	json
// @Produce 	json
// @Param		Authorization	header	string false "Access token with Bearer prefix (optional)"
// @Param 		page_number 	query 	int false "Page number"
// @Param		lower_price		query	int false "Lower price limit"
// @Param		higher_price	query	int false "Higher price limit"
// @Param 		order_by		query	string false "Order option (\"creation_time\" or \"price\")"
// @Param		ascending		query	bool false "Ascending or descending order"
// @Success 	200 {object} 			types.GetFeedResponse "Successfully got feed"
// @Failure 	400 {string} 			string "Bad request"
// @Failure 	401 {string} 			string "Unauthorized"
// @Failure 	404 {string} 			string "Object not found"
// @Failure 	500 {string} 			string "Internal error"
// @Router		/v1/ads/feed	[get]
func (h *AdHandler) getFeed(w http.ResponseWriter, r *http.Request) {
	req := types.CreateGetFeedRequest(r, h.svcCfg)

	res, err := h.adSvc.GetAdFeed(r.Context(), req.Opts)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
		return
	}

	response.WriteResponse(w, types.GetFeedResponse{Feed: res}, http.StatusOK)
}
