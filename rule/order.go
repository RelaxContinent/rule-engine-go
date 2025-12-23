package rule

import "time"

// Order は注文情報を表す
type Order struct {
	ID        string
	UserID    string
	Amount    int
	ItemCount int
	CreatedAt time.Time
	IsPremium bool
	Category  string
}

// User はユーザー情報を表す
type User struct {
	ID              string
	Name            string
	Age             int
	MembershipLevel string // "bronze", "silver", "gold", "platinum"
	TotalPurchases  int
	IsVerified      bool
}

// Result は判定結果を表す
type Result struct {
	Approved     bool
	DiscountRate int
	Rank         string
	Reason       string
}
