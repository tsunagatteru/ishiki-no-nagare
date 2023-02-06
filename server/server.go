package server

import(
	"log"
	"net/http"
	"database/sql"
	"strconv"
	"io/ioutil"
	"encoding/json"
	
	"github.com/gin-gonic/gin"
	"github.com/tsunagatteru/ishiki-no-nagare/db"
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


func showPosts(c *gin.Context){
	page := "0"
	page = (c.Param("page"))
	res, err := http.Get(c.MustGet("BaseURL").(string) + c.MustGet("ApiDir").(string) + "/posts/" + page)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Println(err)
	}
	var response []model.Post
	json.Unmarshal(body, &response)
	c.HTML(http.StatusOK, "posts.tmpl", gin.H{
		"Posts": response,
	})
}

func getPosts(c *gin.Context){
	pageNumber := 0
	pageNumber, err := strconv.Atoi(c.Param("page"))
	if err != nil{
		log.Println(err)
	}
	pageLength := c.MustGet("PageLength").(int)
	dbConn := c.MustGet("dbConn").(*sql.DB)
	posts := db.RetrievePage(dbConn, pageNumber, pageLength)
	c.JSON(http.StatusOK, posts)
}

func createPost(c *gin.Context){
	messageByte, err := c.GetRawData()
	if err != nil {
		log.Println(err)
	}
	message := string(messageByte)
	dbConn := c.MustGet("dbConn").(*sql.DB)
	db.AddPost(dbConn, message)
	c.IndentedJSON(http.StatusCreated, message)
}
