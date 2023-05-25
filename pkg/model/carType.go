package model

type CarType struct {
	TypeID  string `json:"typeID" gorm:"type:bigint unsigned;comment:车辆型号ID;primarykey;autoIncrement;not null;"`
	CarType string `json:"carType" gorm:"type:varchar(10);comment:车辆型号;"`
	Model
}
