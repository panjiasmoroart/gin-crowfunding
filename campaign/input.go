package campaign

import "gin-crowfunding/user"

type GetCampaignDetailInput struct {
	// ada beberapa cara mengirim parameter, json, query params dan uri
	ID int `uri:"id" binding:"required"`
}

// untuk user kita tidak ambil datanya dari json melainkan middleware
type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

type CreateCampaignImageInput struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary" binding:"required"`
}
