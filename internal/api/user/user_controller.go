package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"picnshop/internal/domain/user"
	jwtHelper "picnshop/pkg/jwt"
	"time"
)

type Controller struct {
	userService *user.Service
}

var secret = "secret"

func NewUserController(service *user.Service) *Controller {
	return &Controller{
		userService: service,
	}
}

//TODO: add password match middleware
func (c *Controller) CreateUser(g *gin.Context) {
	var req CreateUserRequest
	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		g.Abort()
		return
	}
	newUser := user.NewUser(req.Username, req.Password)
	err := c.userService.Create(newUser)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusCreated, CreateUserResponse{
		Username: req.Username,
	})
}

func (c *Controller) Login(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
	}
	currentUser, err := c.userService.GetUser(req.Username, req.Password)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   currentUser.ID,
		"username": currentUser.Username,
		"iat":      time.Now().Unix(),
		"iss":      os.Getenv("ENV"),
		"exp": time.Now().Add(24 *
			time.Hour).Unix(),
		//"roles": user.Roles,
	})
	//TODO: get secret from config
	//TODO: update user
	token := jwtHelper.GenerateToken(jwtClaims, secret)

	g.JSON(http.StatusOK, token)
}

func (c *Controller) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, secret, os.Getenv("ENV"))

	g.JSON(http.StatusOK, decodedClaims)

}
