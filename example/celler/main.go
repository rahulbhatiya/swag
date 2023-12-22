package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/controller"
	_ "github.com/swaggo/swag/example/celler/docs"
	"github.com/swaggo/swag/example/celler/httputil"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Release Bundles API
//	@version		1.0
//	@description	APIs for Listing and Deleting the Artifactory Release Bundles.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		ListReleaseBundles := v1.Group("/ListReleaseBundles")
		{
			ListReleaseBundles.GET("", c.ListReleaseBundles)
		}

		VerDeleteReleaseBundles := v1.Group("/VerDeleteReleaseBundles")
		{
			VerDeleteReleaseBundles.DELETE("", c.VerDeleteReleaseBundles)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}
