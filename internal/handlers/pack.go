package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"packer/internal/model"
)

func PackHandler(c *gin.Context) {
	var packs model.Packs
	// Check input ie. new object attributes from request body.
	if err := c.ShouldBindJSON(&packs); err != nil {
		// Handle error in request body.
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	rval := []int{}
	r := map[string]interface{}{
		"Status": "Ok",
		"Object": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}
