package usecases

import (
	"context"
	"marketplace/internal/domain"
)

type AdService interface {
	CreateAd(ctx context.Context, ad *domain.Ad) (*domain.Ad, error)
	GetAdFeed(ctx context.Context, opts domain.GetAdsOpts) ([]domain.FeedPageItem, error)
}
