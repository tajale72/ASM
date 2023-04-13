package router

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"myapi/controller"
)

const (
	Authorization = "authorization"
)

type Process struct {
	r    Router
	auth *gin.RouterGroup
}
type Router struct {
	logger        *log.Logger
	controllersvc controller.ControllerInterface
}

// Application starts
// Loads of the router, controller and database services
func Start() {
	log.Println("I have started")
	var p Process
	p.r.logger = log.New(os.Stdout, servicename, log.LstdFlags|log.Lshortfile)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")
	//router.Use(p.JWTMiddleware)
	p.auth = router.Group("/v1", p.JWTMiddleware)

	p.r.InitTokenEndpoints(router)

	p.r.InitEndpoints(p.auth)
	//Inisitailizing the router's and creating end points
	//Storing the controller, database service in the router to acll the controller and database services.
	p.r.InitializeServices()

	router.Run()

}

// JWTMiddleware is take all the API requests and validates the jwt token.
func (p *Process) JWTMiddleware(c *gin.Context) {

	token := c.Request.Header.Get(Authorization)

	if len(token) == 0 {
		err := errors.New("error from token validation")
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error()})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else {
		updatedtoken := strings.Split(token, " ")
		_, err := p.r.controllersvc.ProcessTokenBody(updatedtoken[1])
		if err != nil {
			log.Println("error from processing the token", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}

}

func (p *Process) Test(token string) bool {
	if len(token) == 0 {
		log.Println("error from token validation")
		return false
	}
	updatedtoken := strings.Split(token, " ")
	log.Println("why am i here", updatedtoken[1])
	_, err := p.r.controllersvc.ProcessTokenBody(updatedtoken[1])
	if err != nil {
		log.Println("error from processing the token", err.Error())
	}
	return true
}
