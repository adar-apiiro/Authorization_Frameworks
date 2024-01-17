package main

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
)

func main() {
	// load the casbin model and policy from files, database is also supported.
	e := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")

	// define your router, and use the Casbin authz middleware.
	// the access that is denied by authz will return HTTP 403 error.
	router := gin.New()
	router.Use(authz.NewAuthorizer(e))

	// Define a route with authorization middleware
	router.GET("/dataset1/resource1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access Granted for GET on /dataset1/resource1"})
	})

	// Define a route with specific permissions required
	router.GET("/dataset1/resource2", authz.Authorize("alice"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access Granted for GET on /dataset1/resource2"})
	})

	// Define a route with a custom unauthorized handler
	router.GET("/dataset2/resource1", authz.Authorize("bob"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access Granted for GET on /dataset2/resource1"})
	})

	// Define a route with a custom unauthorized handler
	router.GET("/dataset2/resource2", authz.Authorize("cathy"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access Granted for GET on /dataset2/resource2"})
	})

	router.Run(":8080")
}
