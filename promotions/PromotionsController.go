package promotions

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgxpool"
	"net/http"
	"strconv"
)

func Controller(route *gin.Engine, db *pgxpool.Pool) {
	promotions := route.Group("/promotions")

	promotions.POST("/csv", func(c *gin.Context) {
		formFile, _, err := c.Request.FormFile("csv-file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		go SaveCSVUseCase(db, formFile)
		c.JSON(http.StatusOK, "In progress")
	})

	promotions.GET("/:id", func(c *gin.Context) {
		idParam, _ := c.Params.Get("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			return
		}
		promotion, err := GetUseCase(db, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, promotion)
	})
}
