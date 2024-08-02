package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello orld"
	c.JSON(http.StatusOK, resp)
}
