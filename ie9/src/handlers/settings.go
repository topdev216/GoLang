package handlers

import (
	"./src/models"
	"github.com/gin-gonic/gin"
)

func GetSettingByKey(key string) string {
	return models.GetSettingByKey(key)
}
