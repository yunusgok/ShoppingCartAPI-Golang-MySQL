package api_helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(g *gin.Context, err error) {

	g.JSON(
		http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
	g.Abort()
	return

}
