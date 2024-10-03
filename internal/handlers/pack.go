package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"packer/internal/logging"
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

    logging.Logger.Debug("Request", slog.String("body", fmt.Sprintf("%+v", packs)))

	if packs.Buckets == nil || len(packs.Buckets) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No buckets provided"})
		return
	}

	// Pack it with buckets
	rval := packer.Pack(packs.Value, packs.Buckets)
	r := map[string]interface{}{
		"Status": "Ok",
		"Packs": rval,
	}

    logging.Logger.Debug("Response", slog.String("body", fmt.Sprintf("%+v", rval)))

	// Finish request nadling
	c.JSON(http.StatusOK, r)
}
