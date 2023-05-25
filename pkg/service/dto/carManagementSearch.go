package dto

type CarManagementSearch struct {
	VinCode        string `json:"vinCode"`        //vin码
	CarNumber      string `json:"carNumber"`      //车牌号
	LeaseCompany   string `json:"leaseCompany"`   //租赁公司
	OnlineDayBegin string `json:"onlineDayBegin"` //日期开始
	OnlineDayEnd   string `json:"onlineDayEnd"`   //日期结束
	CarStatus      string `json:"carStatus"`      //状态
}
