package model

type InsuranceType struct {
	TypeID        string `json:"typeID" gorm:"type:bigint unsigned;comment:险种ID;primarykey;autoIncrement;not null;"`
	InsuranceType string `json:"insuranceType" gorm:"type:varchar(10);comment:险种;"`
	Model
}
