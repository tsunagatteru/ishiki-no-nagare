package server

import (
	"database/sql"
	"embed"
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

func RunRouter(r *gin.Engine, dbConn *sql.DB, config *model.Config, resources embed.FS) {
	r.SetHTMLTemplate(template.Must(template.New("").ParseFS(resources, "res/templates/*.tmpl")))
	resRoot, err := fs.Sub(resources, "res")
	if err != nil {
		log.Fatalln(err)
	}
	staticRoot, err := fs.Sub(resRoot, "static")
	if err != nil {
		log.Fatalln(err)
	}
	r.StaticFS("/static", http.FS(staticRoot))
	r.Use(DatabaseMiddleware(dbConn))
	r.Use(ConfigMiddleware(config))
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte(config.CookieKey))))

	admin := r.Group("/admin")
	admin.Use(AuthRequired)
	admin.GET("/status", status)

	api := r.Group(config.ApiDir)
	api.GET("/posts/:page", getPosts)
	api.POST("/create-post", createPost)

	r.GET("/posts/:page", showPosts)
	r.GET("/", indexRedirect)
	r.POST("/login", login)
	r.GET("/logout", logout)
	r.GET("/create-post", showSubmitPage)

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
		c.Set("ApiDir", config.ApiDir)
		c.Set("PageLength", config.PageLength)
		c.Set("Username", config.UserName)
		c.Set("Password", config.Password)
	}
}
