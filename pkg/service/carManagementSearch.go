package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gokuovo/project-layout/configs"
	"github.com/gokuovo/project-layout/pkg/model"
	"github.com/gokuovo/project-layout/pkg/response"
	"github.com/gokuovo/project-layout/pkg/service/dto"
	"gorm.io/gorm"
)

// SearchAllCar 查看全部车辆信息
func SearchAllCar(ctx *gin.Context) {
	db := configs.GetDB()
	//创建用户
	var newCar []model.CarManagement
	result := db.Find(&newCar)
	if result.Error == nil {
		response.Success(ctx, "查找成功", gin.H{
			"body": newCar,
		})
	} else {
		response.Success(ctx, "查找失败", nil)
	}

}

// SearchCar 根据条件查看车辆信息
func SearchCar(ctx *gin.Context) {
	//db := configs.GetDB()
	//创建用户
	var search dto.CarManagementSearch
	//var car model.CarManagement
	err := ctx.ShouldBindJSON(&search)
	if err != nil {
		panic("接受数据出错，err：" + err.Error())
	}

	////创建车辆
	//newSearch := dto.CarManagementSearch{
	//	VinCode        :search.VinCode,
	//	CarNumber      :search.CarNumber,
	//	LeaseCompany   :search.LeaseCompany,
	//	OnlineDayBegin :search.OnlineDayBegin ,
	//	OnlineDayEnd   :search.OnlineDayEnd,
	//	CarStatus      :search.CarStatus,
	//}

	//result := db.Select("CarID", "VinCode", "Driver", "CarNumber", "LeaseCompany", "CarType",
	//	"CarColor", "OnlineDay", "InsuranceStatus", "InsuranceType", "CarStatus", "OperationalArea",
	//	"PositionMonitor", "CarImageUrl").
	//	Where(map[string]interface{}{
	//		"vin_code": search.VinCode,
	//		"car_number":    search.CarNumber,
	//		"lease_company": search.LeaseCompany,
	//		//"online_day BETWEEN": search.OnlineDayBegin,
	//		//"":                   search.OnlineDayEnd,
	//		"car_status": search.CarStatus,
	//	}).
	//	Take(&car)

	//result := db.Scopes(VinCodeCar,CarNumberCar,LeaseCompanyCar,OnlineDayCar,CarStatusCar).Find(&car)

	//if result.Error == nil {
	//	response.Success(ctx, "查找成功", gin.H{
	//		"body": car,
	//	})
	//} else {
	//	response.Success(ctx, "查找失败", nil)
	//}

}

func VinCodeCar(db *gorm.DB, search dto.CarManagementSearch) *gorm.DB {
	if len(search.VinCode) != 0 {
		return db.Where("vin_code = ?", search.VinCode)
	} else {
		return nil
	}
}

func CarNumberCar(db *gorm.DB, search dto.CarManagementSearch) *gorm.DB {
	if len(search.CarNumber) != 0 {
		return db.Where("car_number = ?", search.CarNumber)
	} else {
		return nil
	}
}

func LeaseCompanyCar(db *gorm.DB, search dto.CarManagementSearch) *gorm.DB {
	if len(search.LeaseCompany) != 0 {
		return db.Where("lease_company = ?", search.LeaseCompany)
	} else {
		return nil
	}
}

func OnlineDayCar(db *gorm.DB, search dto.CarManagementSearch) *gorm.DB {
	if len(search.OnlineDayBegin) != 0 && len(search.OnlineDayEnd) != 0 {
		return db.Where("online_day between ? and ?", search.OnlineDayBegin, search.OnlineDayEnd)
	} else {
		return nil
	}
}

func CarStatusCar(db *gorm.DB, search dto.CarManagementSearch) *gorm.DB {
	if len(search.CarStatus) != 0 {
		return db.Where("car_status = ?", search.CarStatus)
	} else {
		return nil
	}
}
