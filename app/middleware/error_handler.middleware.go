package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Cause   string `json:"cause,omitempty"` // The cause of the error for easier debugging
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process the request

		// Check if there are errors in the Gin context
		if len(c.Errors) > 0 {
			// Only handle the first error for simplicity
			err := c.Errors[0]

			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "An internal server error occurred.",
				Cause:   err.Error(), // Include error details; adjust for production
			})
		}
	}
}
