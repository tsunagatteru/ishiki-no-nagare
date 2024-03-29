package server

import (
	"database/sql"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tsunagatteru/ishiki-no-nagare/db"
)

func Run(resources fs.FS, variables *viper.Viper, config *viper.Viper) {
	dbConn := db.Open(variables.GetString("data-path"))
	defer dbConn.Close()
	db.CreateTable(dbConn)
	r := gin.New()
	r.SetTrustedProxies(nil)
	r.SetHTMLTemplate(template.Must(template.New("").ParseFS(resources, "templates/*.tmpl")))
	staticRoot, err := fs.Sub(resources, "static")
	if err != nil {
		log.Fatalln(err)
	}
	r.StaticFS("/static", http.FS(staticRoot))
	imagesRoot, err := fs.Sub(os.DirFS(variables.GetString("data-path")), "images")
	if err != nil {
		log.Fatalln(err)
	}
	r.StaticFS("/images", http.FS(imagesRoot))
	r.Use(Middleware(dbConn, variables, config))
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte((config.GetString("cookiekey"))))))
	api := r.Group("/api")
	api.GET("/posts/:page", getPosts)
	api.GET("post/:id", getPost)
	api.POST("/login", login)
	api.POST("/logout", logout)
	admin := api.Group("/admin")
	admin.Use(AuthRequired)
	admin.GET("/status", status)
	admin.POST("/create-post", createPost)
	admin.POST("/config", changeConfig)
	r.GET("/post/:id", showPost)
	r.GET("/posts/:page", showPosts)
	r.GET("/", showIndex)
	r.GET("/admin", showAdminPage)
	r.Run(config.GetString("host") + ":" + config.GetString("port"))
}

func Middleware(dbConn *sql.DB, variables *viper.Viper, config *viper.Viper) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbConn", dbConn)
		c.Set("variables", variables)
		c.Set("config", config)
	}
}
