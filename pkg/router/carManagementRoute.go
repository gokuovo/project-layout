package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gokuovo/project-layout/pkg/middleware"
	"github.com/gokuovo/project-layout/pkg/service"
)

func CollectRoute(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware())
	carManagementGroup := r.Group("/carManagement")
	{
		//新增车辆信息
		carManagementGroup.POST("/addCar", service.AddCar)

		//删除单个车辆信息

		//删除批量车辆信息

		//修改车辆信息
		carManagementGroup.PUT("/modifyCar", service.ModifyCar)

		//查询车辆信息

		////selectone
		//userGroup.GET("/selectOne", service.SelectOneUser)
		////selectall
		//userGroup.GET("/selectAll", service.SelectAllUser)
		////delete
		//userGroup.DELETE("/deleteUser", service.DeleteUser)
		////UnscopedDelete
		//userGroup.DELETE("/unscopedDeleteUser", service.UnscopedDeleteUser)
	}
}
