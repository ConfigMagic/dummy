package handler

import (
	"config_saver/internal/db"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	name := c.Param("name")
	if name == "" {
		h.Logger.Error("Name is required in URL")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required in URL"})
		return
	}
	var config db.Config
	if err := c.BindJSON(&config); err != nil {
		h.Logger.Error("Invalid input", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	config.Name = name
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

func (h *ConfigHandler) PushConfigMultipart(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}
	file, err := c.FormFile("archive")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "archive file is required"})
		return
	}
	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open uploaded file"})
		return
	}
	defer f.Close()
	archiveBytes, err := io.ReadAll(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read archive"})
		return
	}
	doc := bson.M{"name": name, "archive": archiveBytes, "updated": time.Now()}
	_, err = h.DB.ConfigCollection.UpdateOne(
		c,
		bson.M{"name": name},
		bson.M{"$set": doc},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		h.Logger.Error("Failed to save archive to MongoDB", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save archive to db"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "saved to db"})
}

// Новый обработчик для выдачи архива из MongoDB
func (h *ConfigHandler) GetConfigArchive(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}
	var result struct {
		Name    string `bson:"name"`
		Archive []byte `bson:"archive"`
	}
	err := h.DB.ConfigCollection.FindOne(c, bson.M{"name": name}).Decode(&result)
	if err != nil {
		h.Logger.Error("Config archive not found", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "Config archive not found"})
		return
	}
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename="+name+".zip")
	c.Writer.Write(result.Archive)
}
