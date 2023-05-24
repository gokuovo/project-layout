package service

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/configs"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/pkg/model"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 新增车辆信息
func AddCar(ctx *gin.Context) {
	db := configs.GetDB()

	var car model.CarManagement

	err := ctx.ShouldBindJSON(&car)
	if err != nil {
		panic("接受数据出错，err：" + err.Error())
	}

	//数据验证
	if len(car.Driver) == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, "司机不能为空", nil)
		return
	}

	if len(car.VinCode) == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, "vin码不能为空", nil)
		return
	}

	if len(car.CarType) == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, "车辆型号不能为空", nil)
		return
	}

	if len(car.CarNumber) == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, "车牌号不能为空", nil)
		return
	}

	if len(car.InsuranceCompany) == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, "保险公司不能为空", nil)
		return
	}

	if len(car.CarColor) == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, "车辆颜色不能为空", nil)
		return
	}

	if len(car.LeaseCompany) == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, "租赁公司不能为空", nil)
		return
	}

	if len(car.InsuranceType) == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, "险种不能为空", nil)
		return
	}

	//创建车辆
	newCar := model.CarManagement{
		Driver:          car.Driver,
		VinCode:         car.VinCode,
		CarNumber:       car.CarNumber,
		LeaseCompany:    car.LeaseCompany,
		CarType:         car.CarType,
		CarColor:        car.CarColor,
		OnlineDay:       car.OnlineDay,
		InsuranceStatus: car.InsuranceStatus,
		InsuranceType:   car.InsuranceType,
		CarStatus:       car.CarStatus,
	}
	result := db.Create(&newCar)
	if result.Error == nil {
		response.Success(ctx, "新增成功", nil)
		fmt.Println("1")
	} else {
		response.Success(ctx, "新增失败", nil)
	}

}
