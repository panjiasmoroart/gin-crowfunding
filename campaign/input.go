package campaign

import "gin-crowfunding/user"

type GetCampaignDetailInput struct {
	// ada beberapa cara mengirim parameter, json, query params dan uri
	ID int `uri:"id" binding:"required"`
}

// untuk user kita tidak ambil datanya dari json melainkan middleware
type CreateCampaignInput struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	GoalAmount       int    `json:"goal_amount"`
	Perks            string `json:"perks"`
	User             user.User
}
