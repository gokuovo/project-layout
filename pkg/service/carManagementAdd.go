package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gokuovo/project-layout/configs"
	"github.com/gokuovo/project-layout/pkg/model"
	"github.com/gokuovo/project-layout/pkg/response"
	"net/http"
	"time"
)

// AddCar 新增车辆信息
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

	if len(car.InsuranceStatus) == 0 {
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

	//上线时间
	onlineDate := time.Now().Format("2006-01-02")
	car.OnlineDay = onlineDate

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
		CarImageUrl:     car.CarImageUrl,
	}

	//新增数据
	result := db.Create(&newCar)
	if result.Error == nil {
		response.Success(ctx, "新增成功", nil)
	} else {
		response.Success(ctx, "新增失败", nil)
	}

}
