package server

import (
	"fmt"
	"mySocialApp/internal/database"
	"mySocialApp/internal/io"
	"net/http"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (s *Server) FollowUser(c *gin.Context) {
	userId := c.Keys["userId"]
	userToFollow := c.Param("followId")

	newFollow := database.Follow{
		FollowerId: cast.ToInt(userToFollow),
		FolloweeId: cast.ToInt(userId),
	}

	foloweeUser, err := s.db.GetUserById(newFollow.FollowerId)
	if err != nil && err != orm.ErrNoRows {
		fmt.Println("Error in getting user ->", err)
		c.JSON(http.StatusInternalServerError, map[string]string{
			"Error": "Unable to process",
		})
		return
	}

	if err == orm.ErrNoRows || foloweeUser.Id == 0 {
		fmt.Println("No user Found")
		c.JSON(http.StatusNotFound, map[string]string{
			"Error": "No User Exists",
		})
		return
	}

	followExist := s.db.CheckForFollow(newFollow.FollowerId, newFollow.FolloweeId)
	if followExist {
		c.JSON(http.StatusOK, map[string]string{
			"Success": "Already following",
		})
		return
	}

	err = s.db.AddNewFollower(newFollow)
	if err != nil {
		fmt.Println("Error in adding a follower", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"success": "followed successfully",
	})
}

func (s *Server) UnfollowUser(c *gin.Context) {
	userId := c.Keys["userId"]
	userToUnfollow := c.Param("unfollowId")

	newFollow := database.Follow{
		FollowerId: cast.ToInt(userToUnfollow),
		FolloweeId: cast.ToInt(userId),
	}

	followExist := s.db.CheckForFollow(newFollow.FollowerId, newFollow.FolloweeId)
	if !followExist {
		c.JSON(http.StatusNotAcceptable, map[string]string{
			"Success": "You are not following the user",
		})
		return
	}

	err := s.db.UnfollowUser(newFollow.FollowerId, newFollow.FolloweeId)
	if err != nil {
		fmt.Println("Error in adding a follower", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"success": "unfollowed successfully",
	})
}

func (s *Server) GetAllFollowers(c *gin.Context) {
	userId := c.Keys["userId"]

	followers, err := s.db.GetAllFollowersForUser(cast.ToInt(userId))
	if err != nil {
		fmt.Println("Error in getting all followers")
		c.JSON(http.StatusInternalServerError, map[string]string{
			"Error": "Error in getting followers" + err.Error(),
		})
	}

	resp := []io.Followers{}
	for _, v := range followers {
		temp := io.Followers{
			User: io.User{Id: v.FolloweeId},
		}
		resp = append(resp, temp)
	}

	c.JSON(http.StatusOK, resp)

}
