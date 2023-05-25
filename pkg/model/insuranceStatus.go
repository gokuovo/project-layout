package model

type InsuranceStatus struct {
	StatusID        string `json:"statusID" gorm:"type:bigint unsigned;comment:保险状态ID;primarykey;autoIncrement;not null;"`
	InsuranceStatus string `json:"insuranceStatus" gorm:"type:varchar(10);comment:保险状态;"`
	Model
}
