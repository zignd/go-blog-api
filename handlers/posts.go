package handlers

import (
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/pkg/errors"
	"github.com/zignd/go-blog-api/blog"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func GetPosts(c *gin.Context) {
	skip, err := strconv.Atoi(c.Query("skip"))
	if err != nil {
		skip = 0
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	if limit > 50 {
		limit = 50
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		c.JSON(500, NewError(errors.Wrap(err, "failed to estabilish a connection").Error()))
		return
	}
	defer session.Close()

	posts, err := blog.GetPosts(session.DB("blog"), skip, limit)
	if err != nil {
		c.JSON(500, NewError(errors.Wrap(err, "failed to retrieve the requested posts").Error()))
		return
	}
	c.JSON(200, posts)
}

func GetPostsTitleUrl(c *gin.Context) {
	c.Param("title-url")
	c.JSON(200, true)
}

func PostPosts(c *gin.Context) {
	var vm Post
	if ok := BindValid(c, &vm); !ok {
		return
	}

	session, db, err := ConnMongoDB()
	if err != nil {
		c.JSON(500, NewError(errors.Wrap(err, "failed to estabilish a connection").Error()))
		return
	}
	defer session.Close()

	post, err := blog.CreatePost(db, &blog.Post{
		Title:     vm.Title,
		Content:   vm.Content,
		AuthorID:  bson.NewObjectId(), // TODO: take this from the JWT once you implement the authentication and authorization
		CreatedAt: time.Now(),
	})
	if err != nil {
		c.JSON(500, NewError(errors.Wrap(err, "failed to create the new post").Error()))
		return
	}

	c.JSON(200, post)
}
