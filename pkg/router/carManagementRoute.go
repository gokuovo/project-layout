package route

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/pkg/middleware"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/pkg/service"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware())
	carManagementGroup := r.Group("/carManagement")
	{
		//新增车辆信息
		carManagementGroup.POST("/addCar", service.AddCar)
		////add
		//userGroup.POST("/add", service.AddUser)
		////selectone
		//userGroup.GET("/selectOne", service.SelectOneUser)
		////selectall
		//userGroup.GET("/selectAll", service.SelectAllUser)
		////modifyUser
		//userGroup.PUT("/modifyUser", service.ModifyUser)
		////delete
		//userGroup.DELETE("/deleteUser", service.DeleteUser)
		////UnscopedDelete
		//userGroup.DELETE("/unscopedDeleteUser", service.UnscopedDeleteUser)
	}
}
