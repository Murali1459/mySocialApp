package server

import (
	"mySocialApp/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.POST("/users/register", s.RegisterUser)
	r.POST("/login", s.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/profile", s.GetProfile)
	auth.GET("/user/:id", s.GetUserById)
	auth.GET("/post/:id", s.GetPostById)
	auth.GET("/posts/:userId", s.GetPostById)
	auth.GET("/posts", s.GetAllPostsForUserId)
	auth.GET("/followers", s.GetAllFollowers)
	auth.GET("/follow/:followId", s.FollowUser)
	auth.GET("/unfollow/:unfollowId", s.UnfollowUser)

	return r
}
