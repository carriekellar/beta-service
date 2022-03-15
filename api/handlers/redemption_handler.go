package handlers

import (
	"beta_service/api/models"
	"beta_service/db/data_access"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RedemptionHandler struct {
	dbAccess *data_access.DbAccess
}

func NewRedemptionHandler(dbAccess *data_access.DbAccess) (*RedemptionHandler, error) {
	return &RedemptionHandler{dbAccess: dbAccess}, nil
}

// This function returns a set of information about redemption for the connected wallet's AXUs
func (h *RedemptionHandler) HandleGetRedemptionInfo(ctx *gin.Context) {
}

// This function parses the KYC form and creates a new KYC entry
func (h *RedemptionHandler) HandleCreateKYC(ctx *gin.Context) {

	var kyc models.KYC

	err := ctx.ShouldBindJSON(&kyc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Succesfully read user info from JSON"})
	}

	err = h.dbAccess.CreateKYC(kyc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Succesfully added KYC information to database"})
	}

	ctx.JSON(http.StatusOK, nil)
}