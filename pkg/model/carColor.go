package model

type CarColor struct {
	ColorID  string `json:"colorID" gorm:"type:bigint unsigned;comment:车辆颜色ID;primarykey;autoIncrement;not null;"`
	CarColor string `json:"carColor" gorm:"type:varchar(10);comment:车辆颜色;"`
	Model
}
