package server

import(
	"log"
	"net/http"
	"database/sql"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/tsunagatteru/ishiki-no-nagare/db"
	_ "github.com/tsunagatteru/ishiki-no-nagare/model"
)

func NewRouter()(*gin.Engine){
	router := gin.New()
	return router
}

func RunRouter(router *gin.Engine, dbConn *sql.DB){
	router.Use(DatabaseMiddleware(dbConn))
	api := router.Group("/api")
	api.GET("/posts/:page", getPosts)
	api.POST("/create-post", createPost)
	router.Run("localhost:8080")
}

func DatabaseMiddleware(dbConn *sql.DB) gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Set("dbConn", dbConn)
	}
}

func getPosts(c *gin.Context){
	pageNumber, err := strconv.Atoi(c.Param("page"))
	if err != nil{
		log.Println(err)
	}
	dbConn := c.MustGet("dbConn").(*sql.DB)
	posts := db.RetrievePage(dbConn, pageNumber)
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
