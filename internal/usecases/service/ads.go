package service

import (
	"context"
	"fmt"
	"marketplace/internal/domain"
	"marketplace/internal/repository"
	"time"
)

type AdService struct {
	adRepo repository.AdRepo
}

func NewAdService(adRepo repository.AdRepo) *AdService {
	return &AdService{
		adRepo: adRepo,
	}
}

func (s *AdService) CreateAd(ctx context.Context, ad *domain.Ad) (*domain.Ad, error) {
	const op = "AdService.CreateAd"

	id, err := s.adRepo.PostAd(ctx, ad)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	ad.Id = id
	ad.CreationTime = time.Now()

	return ad, nil
}

func (s *AdService) GetAdFeed(ctx context.Context, opts domain.GetAdsOpts) ([]domain.FeedPageItem, error) {
	const op = "AdService.GetAdFeed"

	feed, err := s.adRepo.GetAdFeedWithOpts(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return feed, nil
}
