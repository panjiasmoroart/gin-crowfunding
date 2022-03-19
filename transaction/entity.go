package transaction

import (
	"gin-crowfunding/user"
	"time"
)

// Representasi dari table transactions
type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	// relationship
	User      user.User
	CreatedAt time.Time
	UpdatedAt time.Time
}
