// Configures how packages insert routes into the system

package system

import (
	. "codenex.us/ralph/podcast-host"
	"errors"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
)

type SystemRouterService interface {
	AddRoute(method string, route string, handler gin.HandlerFunc) error
	AddRoutes()
}

func (s *System) AddRoute(method string, route string, handler gin.HandlerFunc) error {
	switch method {
	case "POST", "post":
		s.Server.POST(route, handler)
	case "GET", "get":
		s.Server.GET(route, handler)
	case "PUT", "put":
		s.Server.PUT(route, handler)
	case "DELETE", "delete":
		s.Server.DELETE(route, handler)
	case "PATCH", "patch":
		s.Server.PATCH(route, handler)
	case "HEAD", "head":
		s.Server.HEAD(route, handler)
	case "OPTIONS", "options":
		s.Server.OPTIONS(route, handler)
	default:
		return errors.New("ERROR: Can't determine method for route " + route)
	}

	return nil
}

func (s *System) AddRoutes() {
	s.AddRoute("GET", "/", Home())
	s.AddRoute("GET", "/ping", Ping())
	s.AddRoute("GET", "/user/get/:id", s.GetUser())
	s.AddRoute("GET", "/user/create", s.CreateUser())
}

/*************************
********          ********
********  ROUTES  ********
********          ********
*************************/

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	}
}

func (s *System) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := s.DB.GetUser(c.Param("id"))

		if err != nil {
			log.Fatal(err)
		}

		if u == nil {
			c.String(http.StatusNotFound, "That user does not exist")
		} else {
			c.String(http.StatusOK, "That user does exist")
		}

	}
}

func (s *System) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		s.DB.CreateUser(&User{FirstName: "Rafael", LastName: "Martinez", Active: true})

		c.String(http.StatusOK, "Hello World!")
	}
}

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
}
