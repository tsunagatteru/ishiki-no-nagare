package server

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tsunagatteru/ishiki-no-nagare/db"
	"github.com/tsunagatteru/ishiki-no-nagare/model"
)

func indexRedirect(c *gin.Context) {
	location := url.URL{Path: "/posts/0"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func showPosts(c *gin.Context) {
	page := "0"
	page = (c.Param("page"))
	res, err := http.Get(c.MustGet("BaseURL").(string) + c.MustGet("ApiDir").(string) + "/posts/" + page)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var response []model.Post
	json.Unmarshal(body, &response)
	c.HTML(http.StatusOK, "posts.tmpl", gin.H{
		"Posts": response,
	})
}

func showSubmitPage(c *gin.Context) {
	c.HTML(http.StatusOK, "create-post.tmpl", gin.H{})
}

func showLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func getPosts(c *gin.Context) {
	pageNumber := 0
	pageNumber, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		log.Println(err)
	}
	pageLength := c.MustGet("PageLength").(int)
	dbConn := c.MustGet("dbConn").(*sql.DB)
	posts := db.RetrievePage(dbConn, pageNumber, pageLength)
	c.JSON(http.StatusOK, posts)
}

func createPost(c *gin.Context) {
	message := c.PostForm("message")
	dbConn := c.MustGet("dbConn").(*sql.DB)
	db.AddPost(dbConn, message)
	c.IndentedJSON(http.StatusCreated, message)
}
