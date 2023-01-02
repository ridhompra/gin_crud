package controllers

import (
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/ridhompra/models"
	"golang.org/x/crypto/bcrypt"
)

func createJWTToken(user models.Users) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = exp
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}
func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. %+\-]+@[a-z0-9.%+\-]+\.[a-z0-9.%+\-]`)
	return Re.MatchString(email)
}
func SignUp(c *gin.Context) {
	var user models.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if len(string(user.Password)) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password must min 6 character",
		})
		return
	}
	if !ValidateEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email not valid",
		})
		return
	}
	hashedpass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedpass)
	models.DB.Where("email=?", user.Email).Find(&user)
	if user.Id != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exist",
		})
		return
	}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"Email":     user.Email,
		"username":  user.Username,
		"message":   "Sign up successfully",
		"user_test": user,
	})
}

type LoginRequest struct {
	Email    string
	Password string
}

func Login(c *gin.Context) {
	req := new(LoginRequest)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, "invalid login credentials")
		return
	}

	users := new(models.Users)
	models.DB.Where("email = ?", req.Email).First(&users)

	if users.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Login Credential",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(req.Password)); err != nil {
		log.Println(err)
	}

	token, exp, err := createJWTToken(*users)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{
		"token": token,
		"exp":   exp,
		"Users": users,
	})
}
