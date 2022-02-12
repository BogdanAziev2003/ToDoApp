package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


//нормально разобраться с этим
func (h *Handler) signUp(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message":"Hello World!",
	})
}