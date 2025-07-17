package usecases

import (
	"context"
	"marketplace/internal/domain"

	"github.com/google/uuid"
)

type AdService interface {
	CreateAd(ctx context.Context, ad *domain.Ad) (uuid.UUID, error)
	GetAdFeed(ctx context.Context, opts domain.GetAdsOpts) ([]domain.FeedPageItem, error)
}
