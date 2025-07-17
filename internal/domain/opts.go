package domain

type GetAdsOpts struct {
	PageNumber int
	LowerPrice, HigherPrice AdPrice
	OrderOption int
	Ascending bool
	UserLogin string
}
