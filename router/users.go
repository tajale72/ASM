package router

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) InsertUser(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	}
	id, err := r.controllersvc.InsertAllUsers(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSONP(http.StatusAccepted, id)
	}
}

func (r *Router) UpdateUserById(c *gin.Context) {
	id := c.Param("id")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	}
	res, err := r.controllersvc.UpdateUserById(body, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSONP(http.StatusAccepted, res)
	}
}

func (r *Router) GetUser(c *gin.Context) {
	res, err := r.controllersvc.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSONP(http.StatusAccepted, res)
	}
}

func (r *Router) GetUserById(c *gin.Context) {
	id := c.Param("id")
	res, err := r.controllersvc.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSONP(http.StatusAccepted, res)
	}

}

func (r *Router) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	res, err := r.controllersvc.DeleteUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSONP(http.StatusAccepted, res)
	}
}
