package model

type LeaseCompany struct {
	LeaseID      string `json:"leaseID" gorm:"type:bigint unsigned;comment:租赁公司ID;primarykey;autoIncrement;not null;"`
	LeaseCompany string `json:"leaseCompany" gorm:"type:varchar(50);comment:租赁公司;"`
	Model
}
