// Configures how packages insert routes into the system

package system

import (
	. "codenex.us/ralph/podcast-host"
	"codenex.us/ralph/podcast-host/system/files"
	"codenex.us/ralph/podcast-host/system/minio"
	"codenex.us/ralph/podcast-host/system/payload"
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
	case "ANY", "any":
		s.Server.Any(route, handler)
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
	s.AddRoute("ANY", "/episode/upload", s.EpisodeUpload())
	s.AddRoute("GET", "/episode/create", s.EpisodeCreate())
}

/*************************
********          ********
********  ROUTES  ********
********          ********
*************************/

func (s *System) Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Payload
		pl := c.MustGet("PL").(*payload.Payload)

		c.HTML(http.StatusOK, "base", pl.Data)
	}
}

func (s *System) EpisodeCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Payload
		pl := c.MustGet("PL").(*payload.Payload)
		pl.Data["View"] = s.Views["EpisodeCreate"]

		c.HTML(http.StatusOK, "base", pl.Data)
	}
}

func (s *System) EpisodeUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Payload
		pl := c.MustGet("PL").(*payload.Payload)
		pl.Data["View"] = s.Views["EpisodeUpload"]

		af, err := files.NewArangoFiles(s.DB)
		if err != nil {
			log.Fatal(err)
		}

		file, obj, err := c.Request.FormFile("episode")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("About to upload")

		af.Upload(obj.Filename,
			file,
			s.Minios[minio.Types["live"]])

	}
}

func (s *System) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := s.DB.GetUser(c.Param("id"))
		if err != nil {
			log.Fatal(err)
		}

		// Payload
		pl := c.MustGet("PL").(*payload.Payload)

		pl.Data["View"] = s.Views["User"]

		if u == nil {
			c.HTML(http.StatusNotFound, "base", pl.Data)
		} else {
			pl.Data["User"] = u
			c.HTML(http.StatusOK, "base", pl.Data)
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

func (s *System) PayloadMaker() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if context-dependent payload exists
		_, exists := c.Get("PL")
		if exists == false {
			pl, err := payload.New(s.Conf)
			if err != nil {
				log.Fatal("Could not create payload")
			}
			c.Set("PL", pl)
		}

		c.Next()
	}
}

func (s *System) PayloadClearer() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Payload
		pl := c.MustGet("PL").(*payload.Payload)
		_ = pl.Clear()
	}
}

func (s *System) ViewChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Payload
		pl := c.MustGet("PL").(*payload.Payload)
		_, ok := pl.Data["View"]
		if ok == false {
			log.Println("No view found")
			pl.Data["View"] = s.Views["Home"]
		}
		c.Next()
	}
}
