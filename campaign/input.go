package campaign

type GetCampaignDetailInput struct {
	// ada beberapa cara mengirim parameter, json, query params dan uri
	ID int `uri:"id" binding:"required"`
}
