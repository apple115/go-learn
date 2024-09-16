package util

import (
	setting "go-gin-example/pkg"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetPage ...
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
