package server

import (
	"fmt"
	"mySocialApp/internal/database"
	"mySocialApp/internal/io"
	"mySocialApp/internal/util"
	"net/http"
	"os"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
)

func (s *Server) GetUserById(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	fmt.Println(c.Value("userId"))

	user, err := s.db.GetUserById(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := io.User{
		Id:         user.Id,
		UserName:   user.UserName,
		Email:      user.Email,
		ProfilePic: user.ProfilePic,
		Bio:        user.Bio,
	}

	c.JSON(http.StatusOK, resp)
}

func (s *Server) RegisterUser(c *gin.Context) {
	user := &io.RegisterUser{}

	err := c.Bind(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	hashedPw, err := util.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	userDb := database.User{
		UserName: user.UserName,
		Email:    user.Email,
		Password: hashedPw,
	}

	existingUser, err := s.db.GetUserByName(user.UserName)
	if err != nil && err != orm.ErrNoRows {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if existingUser.Id > 0 {
		c.JSON(http.StatusOK, map[string]string{"Error": "Username already present"})
		return
	}

	err = s.db.AddUser(&userDb)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"Success": "User Successfully registered"})
}

func (s *Server) Login(c *gin.Context) {
	req := &io.Login{}
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	user, err := s.db.GetUserByName(req.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if !util.IsSamePassword(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, map[string]string{"Error": "Unauthorized"})
		return
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})

}

func (s *Server) GenerateToken(user database.User) (string, error) {
	secret := cast.ToString(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(1 * time.Hour).Unix(),
		"iat":    time.Now().Unix(),
		"userId": user.Id,
	})

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *Server) GetProfile(c *gin.Context) {
	id := cast.ToInt(c.Keys["userId"])

	user, err := s.db.GetUserById(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := io.User{
		Id:         user.Id,
		UserName:   user.UserName,
		Email:      user.Email,
		ProfilePic: user.ProfilePic,
		Bio:        user.Bio,
	}

	c.JSON(http.StatusOK, resp)
}
