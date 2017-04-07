// Configures the complete server (http, fs, database, etc)
package system

import (
	"codenex.us/ralph/podcast-host/system/db"
	"codenex.us/ralph/podcast-host/system/files"
	"codenex.us/ralph/podcast-host/system/godo"
	"codenex.us/ralph/podcast-host/system/minio"
	"codenex.us/ralph/podcast-host/system/payload"
	"codenex.us/ralph/podcast-host/system/view"
	"github.com/fvbock/endless"
	"github.com/microcosm-cc/bluemonday"
	"github.com/olahol/melody"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
	"html/template"
	"log"
	"os"
)

// Group routing, websockets, db and config
type System struct {
	Server     *gin.Engine
	WS         *melody.Melody
	DB         *db.ArangoDB
	Sanitizers map[string]*bluemonday.Policy
	Conf       *viper.Viper
	PL         *payload.Payload
	Views      map[string]*view.View
	DO         *godo.GoDO
	Minios     map[string]*minio.Minio
	TusC       *files.TusClient
}

// Creates the system
func New() (*System, error) {
	s := new(System)

	err := s.init()

	if err != nil {
		return nil, err
	}

	return s, nil
}

// Runs the system with graceful restarts
func (s *System) Start() {
	err := endless.ListenAndServe(":"+s.Conf.GetString(s.Conf.GetString("Env")+"WebPort"), s.Server)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// Setup input sanitizer
func (s *System) SanitizeInit() {
	// Inititate various policies for sanitization of user input
	s.Sanitizers = make(map[string]*bluemonday.Policy)

	// Policy for bodies of text input by users
	s.Sanitizers["body"] = bluemonday.NewPolicy()
	s.Sanitizers["body"].AllowStandardURLs()
	s.Sanitizers["body"].RequireParseableURLs(true)
	s.Sanitizers["body"].RequireNoFollowOnLinks(true)
	s.Sanitizers["body"].AllowAttrs("href").OnElements("a")
	s.Sanitizers["body"].AllowElements("p")
	s.Sanitizers["body"].AllowImages()
	s.Sanitizers["body"].AllowLists()

	// Policy to strip out all HTML
	s.Sanitizers["strict"] = bluemonday.StrictPolicy()
}

// Setup minio servers
func (s *System) MinioInit() error {
	var err error
	var mc *minio.Conf
	s.Minios = make(map[string]*minio.Minio)

	// The Live file server
	mc, err = minio.NewConf(s.Conf.GetString(s.Conf.GetString("Env")+"MIO_LUrl"),
		s.Conf.GetString(s.Conf.GetString("Env")+"MIO_LATkn"),
		s.Conf.GetString(s.Conf.GetString("Env")+"MIO_LSTkn"))
	if err != nil {
		return err
	}
	s.Minios["live"], err = minio.New(mc)

	// The Archive file server
	mc, err = minio.NewConf(s.Conf.GetString(s.Conf.GetString("Env")+"MIO_AUrl"),
		s.Conf.GetString(s.Conf.GetString("Env")+"MIO_AATkn"),
		s.Conf.GetString(s.Conf.GetString("Env")+"MIO_ASTkn"))
	if err != nil {
		return err
	}
	s.Minios["archive"], err = minio.New(mc)

	// The Backup file server
	mc, err = minio.NewConf(s.Conf.GetString(s.Conf.GetString("Env")+"MIO_BUrl"),
		s.Conf.GetString(s.Conf.GetString("Env")+"MIO_BATkn"),
		s.Conf.GetString(s.Conf.GetString("Env")+"MIO_BSTkn"))
	if err != nil {
		return err
	}
	s.Minios["backup"], err = minio.New(mc)

	return nil
}

// Initializes the system
func (s *System) init() error {
	var err error

	/* HTTP Server
	 */
	s.Server = gin.Default()
	s.Server.Use(s.ViewChecker())
	s.Server.Use(s.PayloadClearer())
	// Set templates
	html := template.Must(template.ParseGlob("tmpl/*"))
	s.Server.SetHTMLTemplate(html)

	/* Websocket
	 */
	s.WS = melody.New()

	/* Config
	 */
	err = s.LoadConfig()

	/* Payload
	 */
	s.PL, err = payload.New()
	if err != nil {
		return err
	}
	// Initialize the payload
	err = s.PL.Init(s.Conf)
	if err != nil {
		return err
	}

	/* Sanitizers
	 */
	s.SanitizeInit()
	if err != nil {
		return err
	}

	/* Database
	 */
	s.DB, err = db.NewArangoDB()
	if err != nil {
		return err
	}
	// Initialize the connection
	err = s.DB.InitConnect(s.Conf.GetString("DBURL"),
		s.Conf.GetString(s.Conf.GetString("Env")+"DBName"),
		s.Conf.GetString(s.Conf.GetString("Env")+"DBUser"),
		s.Conf.GetString(s.Conf.GetString("Env")+"DBPass"))
	if err != nil {
		return err
	}
	// Check that the models exist in the database
	err = s.DB.ModelCheck()
	if err != nil {
		return err
	}

	/* Digital Ocean
	 */
	// Initialize the DO struct
	s.DO, err = godo.New(s.Conf.GetString(s.Conf.GetString("Env") + "DOToken"))
	if err != nil {
		return err
	}

	/* Minio
	 */
	err = s.MinioInit()
	if err != nil {
		return err
	}

	/* TUS Client
	 */
	s.TusC, err = files.NewTusClient(s.Conf.GetString(s.Conf.GetString("Env")+"MIO_LUrl"), nil)
	if err != nil {
		return err
	}

	/* Routes
	 */
	s.AddRoutes()

	/* View
	 */
	s.Views, _ = view.Views()

	/* Static Files
	 */
	s.Server.Static("/css", "./css")
	s.Server.Static("/scripts", "./scripts")
	s.Server.Static("/images", "./images")

	return nil
}
