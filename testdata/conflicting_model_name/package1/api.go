package package1

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/testdata/conflicting_model_name/package1/data"
)

// @Description get struct package1.data.Foo
// @ID get-struct-package1-foo
// @Accept  json
// @Produce  json
// @Success 200 {object} data.Foo
// @Router /testapi/get-struct-package1-foo [get]
func GetStructPackage1Foo(c *gin.Context) {
	//write your code
	var _ = data.Foo{}
}
