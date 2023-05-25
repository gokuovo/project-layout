package model

type StatusCode struct {
	StatusID  string `json:"statusID" gorm:"type:bigint unsigned;comment:状态ID;primarykey;autoIncrement;not null;"`
	CarStatus string `json:"carStatus" gorm:"type:varchar(10);comment:状态;"`
	Model
}
