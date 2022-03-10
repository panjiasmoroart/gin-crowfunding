package handler

import (
	"gin-crowfunding/campaign"
	"gin-crowfunding/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// langkah pengerjaan
// tangkap parameter di handler
// handler ke service
// service yg menentukan  repository mana yg di call
// repository : FindAll, FindByUserID  -> akses ke db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// convert dulu string to int
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
