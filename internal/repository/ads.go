package repository

import (
	"context"
	"marketplace/internal/domain"

	"github.com/google/uuid"
)

type AdRepo interface {
	PostAd(ctx context.Context, ad *domain.Ad) (uuid.UUID, error)
	GetAdsWithOpts(ctx context.Context, opts domain.GetAdsOpts) ([]*domain.Ad, error)
}
