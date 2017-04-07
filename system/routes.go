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

// Simplify adding routes
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

// Register routes with system
func (s *System) AddRoutes() {
	s.AddRoute("GET", "/", s.Home())
	s.AddRoute("GET", "/ping", Ping())
	s.AddRoute("GET", "/user/get/:id", s.GetUser())
	s.AddRoute("GET", "/user/create", s.CreateUser())
}

/*************************
********          ********
********  ROUTES  ********
********          ********
*************************/

func (s *System) Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "base", s.PL.Data)
	}
}

func (s *System) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := s.DB.GetUser(c.Param("id"))
		if err != nil {
			log.Fatal(err)
		}

		s.PL.Data["View"] = s.Views["User"]

		if u == nil {
			c.HTML(http.StatusNotFound, "base", s.PL.Data)
		} else {
			s.PL.Data["User"] = u
			c.HTML(http.StatusOK, "base", s.PL.Data)
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

/*****************************
********              ********
********  MIDDLEWARE  ********
********              ********
*****************************/
func (s *System) PayloadClearer() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		_ = s.PL.Clear()
	}
}

func (s *System) ViewChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := s.PL.Data["View"]

		if ok == false {
			log.Println("No view found")
			s.PL.Data["View"] = s.Views["Home"]
		}
		c.Next()
	}
}
