package handler

import (
    "net/http"
    "wallet-service/models"
    "wallet-service/internal/wallet"
    "wallet-service/internal/db"
    "github.com/gin-gonic/gin"
)

func CreateWallet(c *gin.Context) {
    var req models.WalletRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    existingWallet, err := db.GetWalletByUserID(req.UserID)
    if err == nil && existingWallet != nil {
        c.JSON(http.StatusOK, existingWallet)
        return
    }

    address, _, err := wallet.GenerateWalletAddress()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate wallet"})
        return
    }

    wallet := models.WalletResponse{UserID: req.UserID, Address: address}
    if err := db.SaveWallet(wallet); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save wallet"})
        return
    }

    c.JSON(http.StatusOK, wallet)
}

