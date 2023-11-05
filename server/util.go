package server

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// handle error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// handle error custom response
func errorCustomResponse(msg string) gin.H {
	return errorResponse(errors.New(msg))
}
