package handler

import (
	"gin-crowfunding/campaign"
	"gin-crowfunding/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// convert dulu string to int,  ?user_id=1  Query Params
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

// singular buat detail campaign
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// repository : get campaign by id
	// service : inputnya struct input => menangkap id di url
	// handler : mapping id yg di url ke struct input => service
}
