package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"packer/internal/model"
	"packer/pkg/packer"
)

// PackHAndler handles a POST request with packing payload.
func PackHandler(c *gin.Context) {
	var packs model.Packs

	// Check input ie. new object attributes from request body.
	if err := c.ShouldBindJSON(&packs); err != nil {
		// Handle error in request body.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pack it with buckets
	rval := packer.Pack(packs.Value, packs.Buckets)
	r := map[string]interface{}{
		"Status": "Ok",
		"Object": rval,
	}

	// Finish request nadling
	c.JSON(http.StatusOK, r)
}
