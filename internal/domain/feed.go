package domain

type FeedPageItem struct {
	ItemNumber 		int		`json:"item_number"`
	SelfAuthored 	bool	`json:"self_authored"`
	Ad 				*Ad		`json:"ad"`
}
