package controllers

// TODO: Fix time
import (
	"golang-wm-api/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func createToken(userData string) (string, error) {
	var expirationTime = time.Now().Add(time.Hour * 24)
	// fmt.Println("date now: ", expirationTime)
	claims := jwt.MapClaims{
		"username": userData,
		"exp":      expirationTime,
	}
	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT")))

	if err != nil {
		return "", err
	}
	return signedToken, err
}

func (ControllerCollection) Login(c *gin.Context) {
	var admin models.User
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Please input correct json data",
			"error":   err,
		})
		return
	}

	if err := models.DB.First(&admin).Where("username", admin.Username).Where("password", admin.Password).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Server error",
				"error":   err.Error(),
			})
			return
		}
	}

	token, err := createToken(admin.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Fail create a token",
			"error":   err,
		})
		return
	}

	// expiration := time.Hour * 24
	// fmt.Println("date now: ", expirationTime)
	// fmt.Println("date now: ", int(expirationTime.Unix()))
	// fmt.Println("ex now: ", expiration)
	// fmt.Println("ex now: ", int(expiration))

	// substractTime := time.Until(expirationTime)
	var expiration = 24 * 60 * 60
	c.SetCookie("jwt", token, int(expiration), "/", "*", false, true)

	// c.SetCookie("jwt", jwtToken, (2*24*60*60), "", )
	c.JSON(http.StatusOK, gin.H{
		"message": "Login success, welcome " + admin.Username,
	})
}
