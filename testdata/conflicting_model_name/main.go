package conflicting_model_name

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/testdata/conflicting_model_name/package1"
	"github.com/swaggo/swag/testdata/conflicting_model_name/package2"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server
// @termsOfService http://swagger.io/terms/

// @host petstore.swagger.io
// @BasePath /v2

func main() {
	r := gin.New()
	r.GET("/testapi/get-string-by-int/:some_id", package1.GetStructPackage1Foo)
	r.GET("//testapi/get-struct-array-by-string/:some_id", package2.GetStructPackage2Foo)
	r.Run()
}
