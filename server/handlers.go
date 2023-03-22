package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tsunagatteru/ishiki-no-nagare/db"
	"github.com/tsunagatteru/ishiki-no-nagare/model"
)

func index(c *gin.Context) {
	res, err := http.Get(c.MustGet("BaseURL").(string) + "/api/posts/1")
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
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Posts": response,
	})
}

func showPosts(c *gin.Context) {
	var page string
	page = (c.Param("page"))
	res, err := http.Get(c.MustGet("BaseURL").(string) + "/api/posts/" + page)
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

func showPost(c *gin.Context) {
	id := (c.Param("id"))
	res, err := http.Get(c.MustGet("BaseURL").(string) + "/api/post/" + id)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var response model.Post
	json.Unmarshal(body, &response)
	c.HTML(http.StatusOK, "post.tmpl", gin.H{
		"Post": response,
	})
}

func showSubmitPage(c *gin.Context) {
	c.HTML(http.StatusOK, "create-post.tmpl", gin.H{})
}

func showLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}
func showLogout(c *gin.Context) {
	c.HTML(http.StatusOK, "logout.tmpl", gin.H{})
}

func showAdminPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.tmpl", gin.H{})
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

func getPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	dbConn := c.MustGet("dbConn").(*sql.DB)
	post := db.RetrievePost(dbConn, id)
	c.JSON(http.StatusOK, post)
}

func createPost(c *gin.Context) {
	message := c.PostForm("message")
	dbConn := c.MustGet("dbConn").(*sql.DB)

	form, err := c.MultipartForm()
	if err != nil {
		log.Println(err)
	}
	dataPath := c.MustGet("DataPath").(string)
	files := form.File["files"]
	for _, file := range files {
		fileExt := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", time.Now().Unix()) + fileExt
		c.SaveUploadedFile(file, dataPath+"images/"+filename)
	}

	db.AddPost(dbConn, message)
	c.IndentedJSON(http.StatusCreated, "created")
}
