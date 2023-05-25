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

		//查询全部车辆信息
		carManagementGroup.GET("/searchAllCar", service.SearchAllCar)

		//根据条件查询车辆信息
		carManagementGroup.GET("/searchCar", service.SearchCar)

		////delete
		//userGroup.DELETE("/deleteUser", service.DeleteUser)
		////UnscopedDelete
		//userGroup.DELETE("/unscopedDeleteUser", service.UnscopedDeleteUser)
	}
}
