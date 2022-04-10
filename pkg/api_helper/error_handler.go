package api_helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleError return error response to  given context with given error
func HandleError(g *gin.Context, err error) {

	g.JSON(
		http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	g.Abort()
	return

}
