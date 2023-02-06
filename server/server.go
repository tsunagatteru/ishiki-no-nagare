package server

import(
	"database/sql"
	
	"github.com/gin-gonic/gin"
	"github.com/tsunagatteru/ishiki-no-nagare/model"
)

func NewRouter()(*gin.Engine){
	router := gin.New()
	return router
}

func RunRouter(router *gin.Engine, dbConn *sql.DB, config *model.Config){
	router.LoadHTMLGlob("res/templates/*.tmpl")
	router.Use(DatabaseMiddleware(dbConn))
	router.Use(ConfigMiddleware(config))
	api := router.Group(config.ApiDir)
	api.GET("/posts/:page", getPosts)
	api.POST("/create-post", createPost)
	front := router.Group("/")
	front.GET("/posts/:page", showPosts)
	router.Run(config.Host + ":" +config.Port)
}

func DatabaseMiddleware(dbConn *sql.DB) gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Set("dbConn", dbConn)
	}
}

func ConfigMiddleware(config *model.Config) gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Set("BaseURL", "http://" + config.Host + ":" + config.Port)
		c.Set("ApiDir", config.ApiDir)
		c.Set("PageLength", config.PageLength)
	}
}

