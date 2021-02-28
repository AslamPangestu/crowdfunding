package routes

import (
	"crowdfunding/handler"
	"crowdfunding/repository"
	"crowdfunding/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CampaignRoute : Campaign Routing
func CampaignRoute(api *gin.RouterGroup, db *gorm.DB) {
	repository := repository.CampaignRepositoryInit(db)
	service := services.CampaignServiceInit(repository)
	handler := handler.CampaignHandlerInit(service)
	api.GET("/campaigns", handler.GetCampaigns)
}
