package types

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"marketplace/internal/config"
	"marketplace/internal/domain"
	"marketplace/pkg/utils"
	"net/http"
	"strconv"
)

func checkPrice(price int, cfg config.ServiceConfig) bool {
	return price > 0 && price <= cfg.MaxPrice
}

func checkTitle(title string, cfg config.ServiceConfig) bool {
	return len(title) < cfg.MaxTitleLength
}

func checkDescription(desc string, cfg config.ServiceConfig) bool {
	return len(desc) < cfg.MaxDescriptionLength
}

func checkImage(image image.Image, cfg config.ServiceConfig) bool {
	return image.Bounds().Dx() <= cfg.MaxImageSize && image.Bounds().Dy() <= cfg.MaxImageSize
}

// Requests ----------------------------------------------------------------------

type CreateAdRequest struct {
	Ad 	domain.Ad	`json:"ad"` 
}

func MakePostCreateAdRequest(r *http.Request, cfg config.ServiceConfig) (*CreateAdRequest, error) {
	const op = "MakeCreateAdRequest"

	var req CreateAdRequest

	if err := json.NewDecoder(r.Body).Decode(&req.Ad); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if !checkPrice(int(req.Ad.Price), cfg) {
		return nil, fmt.Errorf("%s: %w", op, ErrBadPriceValue)
	} else if !checkTitle(req.Ad.Title, cfg) {
		return nil, fmt.Errorf("%s: %w", op, ErrBadTitleLength)
	} else if !checkDescription(req.Ad.Description, cfg) {
		return nil, fmt.Errorf("%s: %w", op, ErrBadDescriptionLength)
	}

	image, _, err := utils.GetImage(req.Ad.ImageAddress)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if !checkImage(image, cfg) {
		return nil, fmt.Errorf("%s: %w", op, ErrBadImageFormat)
	}

	req.Ad.AuthorLogin = r.Header.Get("X-User-Login")

	return &req, nil
}

type GetFeedRequest struct {
	Opts	domain.GetAdsOpts
}

func CreateGetFeedRequest(r *http.Request, cfg config.ServiceConfig) *GetFeedRequest {
	const op = "CreateGetFeedRequest"

	req := GetFeedRequest{
		Opts: domain.GetAdsOpts{
			PageNumber: 1,
			LowerPrice: 0,
			HigherPrice: domain.AdPrice(cfg.MaxPrice),
			OrderOption: domain.OrderByCreationTime,
			Ascending: false,
			UserLogin: r.Header.Get("X-User-Login"),
		},
	}

	if number, err := strconv.Atoi(r.URL.Query().Get("page_number")); err == nil {
		req.Opts.PageNumber = number
	}

	if price, err := strconv.Atoi(r.URL.Query().Get("lower_price")); err == nil {
		req.Opts.LowerPrice = domain.AdPrice(price)
	}

	if price, err := strconv.Atoi(r.URL.Query().Get("higher_price")); err == nil {
		req.Opts.HigherPrice = domain.AdPrice(price)
	}

	switch r.URL.Query().Get("order_by") {
	case "creation_time":
		req.Opts.OrderOption = domain.OrderByCreationTime
	case "price":
		req.Opts.OrderOption = domain.OrderByPrice
	}

	if asc, err := strconv.ParseBool(r.URL.Query().Get("ascending")); err == nil {
		req.Opts.Ascending = asc
	}

	return &req
}

// Responses ---------------------------------------------------------------------

type PostCreateAdResponse struct {
	AdId 	string					`json:"ad_id"`
}

type GetFeedResponse struct {
	Feed 	[]domain.FeedPageItem	`json:"feed"`
}
