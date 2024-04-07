package controllers

import "github.com/gin-gonic/gin"

// get users
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// get single user
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// sign up
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// login
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// hash password
func HashPassword(password string) string {
	return ""
}

// verify password
func VerifyPassword(userPassword string, previousPassword string) (bool, string) { return "" }
