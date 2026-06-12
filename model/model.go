package model

type Item struct {
	ID       int
	Name     string
	SellerID int
}

type GetItemSeller struct {
	ItemName   string
	SellerName string
}