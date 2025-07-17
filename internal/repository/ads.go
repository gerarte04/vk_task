package repository

import (
	"context"
	"marketplace/internal/domain"

	"github.com/google/uuid"
)

type AdRepo interface {
	PostAd(ctx context.Context, ad *domain.Ad) (uuid.UUID, error)
	GetAdFeedWithOpts(ctx context.Context, opts domain.GetAdsOpts) ([]domain.FeedPageItem, error)
}
