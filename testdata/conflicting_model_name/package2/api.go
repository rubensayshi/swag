package package2

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/testdata/conflicting_model_name/package2/data"
)

// @Description get struct package2.data.Foo
// @ID get-struct-package2-foo
// @Accept  json
// @Produce  json
// @Success 200 {object} data.Foo
// @Router /testapi/get-struct-package2-foo [get]
func GetStructPackage2Foo(c *gin.Context) {
	//write your code
	var _ = data.Foo{}
}
