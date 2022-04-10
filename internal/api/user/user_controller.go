package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"picnshop/internal/config"
	"picnshop/internal/domain/user"
	"picnshop/pkg/api_helper"
	jwtHelper "picnshop/pkg/jwt"
	"strconv"
	"time"
)

type Controller struct {
	userService *user.Service
	appConfig   *config.Configuration
}

func NewUserController(service *user.Service, appConfig *config.Configuration) *Controller {
	return &Controller{
		userService: service,
		appConfig:   appConfig,
	}
}

func (c *Controller) CreateUser(g *gin.Context) {
	var req CreateUserRequest
	if err := g.ShouldBind(&req); err != nil {
		g.JSON(
			http.StatusBadRequest, gin.H{
				"error_message": "Check your request body.",
			})
		g.Abort()
		return
	}
	newUser := user.NewUser(req.Username, req.Password, req.Password2)
	err := c.userService.Create(newUser)
	if err != nil {
		g.JSON(
			http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
		g.Abort()
		return
	}

	g.JSON(
		http.StatusCreated, CreateUserResponse{
			Username: req.Username,
		})
}

func (c *Controller) Login(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		g.JSON(
			http.StatusBadRequest, gin.H{
				"error_message": "Check your request body.",
			})
	}
	currentUser, err := c.userService.GetUser(req.Username, req.Password)
	if err != nil {
		g.JSON(
			http.StatusNotFound, gin.H{
				"error_message": err.Error(),
			})
		return
	}
	decodedClaims := jwtHelper.VerifyToken(currentUser.Token, c.appConfig.SecretKey)
	if decodedClaims == nil {
		jwtClaims := jwt.NewWithClaims(
			jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   strconv.FormatInt(int64(currentUser.ID), 10),
				"username": currentUser.Username,
				"iat":      time.Now().Unix(),
				"iss":      os.Getenv("ENV"),
				"exp": time.Now().Add(
					24 *
						time.Hour).Unix(),
				"isAdmin": currentUser.IsAdmin,
			})
		token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.SecretKey)
		currentUser.Token = token
		err = c.userService.UpdateUser(&currentUser)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
	}

	g.JSON(
		http.StatusOK, LoginResponse{Username: currentUser.Username, UserId: currentUser.ID, Token: currentUser.Token})
}

func (c *Controller) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, c.appConfig.SecretKey)

	g.JSON(http.StatusOK, decodedClaims)

}
