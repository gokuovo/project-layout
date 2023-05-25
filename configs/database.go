package configs

import (
	"fmt"
	models "github.com/gokuovo/project-layout/pkg/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	user := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})

	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	db.AutoMigrate(&models.CarManagement{}, &models.CarColor{}, &models.CarType{},
		&models.InsuranceStatus{}, &models.InsuranceType{}, &models.LeaseCompany{},
		&models.StatusCode{})

	DB = db
	return db

}

func GetDB() *gorm.DB {
	return DB
}
