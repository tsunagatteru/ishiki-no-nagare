package server

import (
	"database/sql"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/tsunagatteru/ishiki-no-nagare/model"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	return router
}

func RunRouter(r *gin.Engine, dbConn *sql.DB, config *model.Config, resources fs.FS) {
	r.SetHTMLTemplate(template.Must(template.New("").ParseFS(resources, "templates/*.tmpl")))
	staticRoot, err := fs.Sub(resources, "static")
	if err != nil {
		log.Fatalln(err)
	}
	r.StaticFS("/static", http.FS(staticRoot))
	r.Use(DatabaseMiddleware(dbConn))
	r.Use(ConfigMiddleware(config))
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte(config.CookieKey))))

	api := r.Group("/api")
	api.GET("/posts/:page", getPosts)
	api.GET("/status", status)
	api.POST("/login", login)
	api.GET("/logout", logout)
	admin := api.Group("/admin")
	admin.Use(AuthRequired)
	admin.POST("/create-post", createPost)

	r.GET("/posts/:page", showPosts)
	r.GET("/", indexRedirect)
	r.GET("/create-post", showSubmitPage)
	r.GET("/login", showLogin)

	r.Run(config.Host + ":" + config.Port)
}

func DatabaseMiddleware(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbConn", dbConn)
	}
}

func ConfigMiddleware(config *model.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("BaseURL", "http://"+config.Host+":"+config.Port)
		c.Set("PageLength", config.PageLength)
		c.Set("Username", config.UserName)
		c.Set("Password", config.Password)
	}
}
