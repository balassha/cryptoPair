package Routes

import (
	"github.com/gin-gonic/gin"

	"cryptoCurrencies/Controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	grp := r.Group("/v1/service")

	{
		grp.GET("/db", Controllers.DBDataController)
		grp.GET("/price", Controllers.GetRawDataController)
	}

	return r
}
