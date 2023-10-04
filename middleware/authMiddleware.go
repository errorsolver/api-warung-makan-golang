package middleware

import (
	"fmt"
	"golang-wm-api/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AdminAuth() gin.HandlerFunc {
	var admin models.User
	return func(c *gin.Context) {
		cookieToken, cookieErr := c.Cookie("jwt")
		if cookieErr != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization failed",
			})
			return
		}

		fmt.Println("token ", cookieToken)
		parseToken, parseErr := jwt.Parse(cookieToken, func(cookieToken *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT")), nil
		})
		fmt.Println("token ", parseToken, "err ", parseErr)

		if !parseToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token not valid",
				"error":   parseErr.Error(),
			})
			return
		}

		claims, ok := parseToken.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Failed to get token claims")
			return
		}

		fmt.Println("claims", claims["username"])

		if err := models.DB.First(&admin).Where("role = ?", "admin").Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization failed",
			})
			return
		}
		c.Next()
	}
}
func UserAuth(c *gin.Context) {
	tokenString, tokenErr := c.Cookie("jwt")
	if tokenErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	}
	parseToken, parseErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, claimsOk := parseToken.Claims.(jwt.MapClaims); claimsOk && parseToken.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
		}
		var user models.User
		if err := models.DB.First(&user).Where("role", "user").Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization failed",
			})
		}
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"error":   parseErr,
		})
	}
}

// func ValidateJWT(next func(w http.ResponseWriter, r* http.Request)) http.Handler{
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

// 		if r.Header["Token"]:=nil{
// 			token, err:= jwt.Parse(r.Header["Token"][0], func(t *jwt.token) (interface{}, error){
// 				_, ok:=t.Method.(*jwt.SigningMethodHMAC)
// 				if !ok {
// 					w.WriteHeader(http.StatusUnauthorized)
// 					w.Write([]byte("not authorized"))
// 				}
// 				return SECRET, nil
// 			})

// 			if err:= nil {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				w.Write([]byte("not authorized"))
// 			}

// 			if token.Valid {
// 				next(w, r)
// 			}
// 		} else{
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("not authorized"))
// 		}
// 	})
// }
