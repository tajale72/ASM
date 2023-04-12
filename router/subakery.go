package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Hello is  a test function
func (r *Router) Hello(c *gin.Context) {
	res, err := r.controllersvc.Hello()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res = "SuBakery"
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": res,
	})

}

func (r *Router) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (r *Router) PostLogin(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)
	log.Println(string(body))

	var finance interface{}
	err := json.Unmarshal(body, &finance)

	if finance == nil {
		log.Println("value nil", finance)
		c.JSON(http.StatusBadRequest, gin.H{"empty payload": finance})
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if err != nil {
		log.Println("errror unmarshalling", err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, finance)
	}
}
