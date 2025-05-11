package handler

import (
	"config_saver/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type ConfigHandler struct {
	DB     *db.MongoDB
	Logger *zap.Logger
}

func (h *ConfigHandler) GetConfig(c *gin.Context) {
	name := c.Param("name")
	var config db.Config
	err := h.DB.ConfigCollection.FindOne(c, bson.M{"name": name}).Decode(&config)
	if err != nil {
		h.Logger.Error("Config not found", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "Config not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": config.Data})
}

func (h *ConfigHandler) SaveConfig(c *gin.Context) {
	var config db.Config
	if err := c.BindJSON(&config); err != nil {
		h.Logger.Error("Invalid input", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	_, err := h.DB.ConfigCollection.InsertOne(c, config)
	if err != nil {
		h.Logger.Error("Failed to save config", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save config"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "saved"})
}

func (h *ConfigHandler) DeleteConfig(c *gin.Context) {
	name := c.Param("name")
	res, err := h.DB.ConfigCollection.DeleteOne(c, bson.M{"name": name})
	if err != nil {
		h.Logger.Error("Failed to delete config", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete config"})
		return
	}
	if res.DeletedCount == 0 {
		h.Logger.Warn("Config not found for deletion", zap.String("name", name))
		c.JSON(http.StatusNotFound, gin.H{"error": "Config not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
