package server

import (
	"fmt"
	"mySocialApp/internal/io"
	"net/http"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (s *Server) GetPostById(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	post, err := s.db.GetPostById(id)
	if err == orm.ErrNoRows {

		c.JSON(http.StatusNotFound, map[string]string{
			"Not Found": "No Post Found",
		})
		return
	}
	if err != nil && err != orm.ErrNoRows {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, map[string]string{
			"Error in getting post": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (s *Server) GetAllPostsForUserId(c *gin.Context) {
	userId := cast.ToInt(c.Keys["userId"])

	posts, err := s.GetAllPosts(userId)
	if err != nil {
		fmt.Println("Error in getting all posts", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (s *Server) GetAllPostsForUser(c *gin.Context) {
	userId := c.Param("userId")

	posts, err := s.GetAllPosts(cast.ToInt(userId))
	if err != nil {
		fmt.Println("Error in getting all posts", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (s *Server) GetAllPosts(userId int) ([]io.Post, error) {
	posts, err := s.db.GetAllPostsForUser(userId)
	if err != nil {
		fmt.Println("Error in getting posts for user -> ", userId)
		return nil, err
	}
	postsResp := []io.Post{}
	for _, v := range posts {
		tmp := io.Post{
			Id:       v.Id,
			UserId:   userId,
			Content:  v.Content,
			ImageUrl: v.ImageUrl,
		}
		postsResp = append(postsResp, tmp)
	}
	return postsResp, nil
}
