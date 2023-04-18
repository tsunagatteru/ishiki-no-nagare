package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tsunagatteru/ishiki-no-nagare/db"
)

func showIndex(c *gin.Context) {
	dbConn := c.MustGet("dbConn").(*sql.DB)
	variables := c.MustGet("variables").(*viper.Viper)
	pageLength := variables.Get("page-length").(int)
	posts := db.RetrievePage(dbConn, 1, pageLength)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Posts": posts,
	})
}

func showPosts(c *gin.Context) {
	pageNumber := 0
	pageNumber, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		log.Println(err)
	}
	dbConn := c.MustGet("dbConn").(*sql.DB)
	variables := c.MustGet("variables").(*viper.Viper)
	pageLength := variables.Get("page-length").(int)
	posts := db.RetrievePage(dbConn, pageNumber, pageLength)

	c.HTML(http.StatusOK, "posts.tmpl", gin.H{
		"Posts": posts,
	})
}

func showPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	dbConn := c.MustGet("dbConn").(*sql.DB)
	post := db.RetrievePost(dbConn, id)
	c.HTML(http.StatusOK, "post.tmpl", gin.H{
		"Post": post,
	})
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
	dbConn := c.MustGet("dbConn").(*sql.DB)
	variables := c.MustGet("variables").(*viper.Viper)
	pageLength := variables.Get("page-length").(int)
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
	if message == "" {
		c.IndentedJSON(http.StatusBadRequest, "empty message")
	} else {
		dbConn := c.MustGet("dbConn").(*sql.DB)
		variables := c.MustGet("variables").(*viper.Viper)
		form, err := c.MultipartForm()
		if err != nil {
			log.Println(err)
		}
		dataPath := variables.Get("data-path").(string)
		files := form.File["files"]
		var filename string
		if len(files) != 0 {
			file := files[0]
			fileExt := filepath.Ext(file.Filename)
			originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
			filename = strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", time.Now().Unix()) + fileExt
			c.SaveUploadedFile(file, dataPath+"images/"+filename)
		} else {
			filename = ""
		}
		db.AddPost(dbConn, message, filename)
		c.IndentedJSON(http.StatusCreated, "created")
	}
}

func changeConfig(c *gin.Context) {
	config := c.MustGet("config").(*viper.Viper)
	username := c.PostForm("username")
	if username != "" {
		config.Set("username", username)
	}
	password := c.PostForm("password")
	if password != "" {
		config.Set("password", password)
	}
	cookiekey := c.PostForm("cookiekey")
	if cookiekey != "" {
		config.Set("cookiekey", cookiekey)
	}
	config.WriteConfig()
	logout(c)
	c.JSON(http.StatusOK, "Config updated")
}
