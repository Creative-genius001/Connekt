package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJobListing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get job listing routes"})
}
