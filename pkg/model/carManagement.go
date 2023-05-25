package model

type CarManagement struct {
	CarID           string `json:"carID" gorm:"type:bigint unsigned;comment:CarID;primarykey;autoIncrement;not null;"`
	VinCode         string `json:"vinCode" gorm:"type:varchar(20);comment:vin码;"`
	Driver          string `json:"driver" gorm:"type:varchar(20);comment:司机;"`
	CarNumber       string `json:"carNumber" gorm:"type:varchar(8);comment:车牌号;"`
	LeaseCompany    string `json:"leaseCompany" gorm:"type:varchar(50);comment:租赁公司;"`
	CarType         string `json:"carType" gorm:"type:varchar(10);comment:车辆型号;"`
	CarColor        string `json:"carColor" gorm:"type:varchar(10);comment:车辆颜色;"`
	OnlineDay       string `json:"onlineDay" gorm:"type:date;comment:上线日期;"`
	InsuranceStatus string `json:"insuranceStatus" gorm:"type:varchar(10);comment:保险状态;"`
	InsuranceType   string `json:"insuranceType" gorm:"type:varchar(10);comment:险种;"`
	CarStatus       string `json:"carStatus" gorm:"type:varchar(10);comment:状态;"`
	OperationalArea string `json:"operationalArea" gorm:"type:varchar(50);comment:运营域;"`
	PositionMonitor string `json:"positionMonitor" gorm:"type:varchar(50);comment:位置监控;"`
	CarImageUrl     string `json:"carImageUrl" gorm:"type:varchar(200);comment:车辆照片;"`
	Model
}
