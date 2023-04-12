package router

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) GenerateToken(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if len(body) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := r.controllersvc.GenerateToken(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, token)
}

func (r *Router) Verify(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(body) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "body is empty"})
	}

	res, err := r.controllersvc.VerifyToken(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusAccepted, res)
	}
}
