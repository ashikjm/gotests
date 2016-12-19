package base

import (
	"github.com/jinzhu/gorm"
	//"fmt"
	//"log"
	//"fmt"
	//"fmt"
	"service/common"
	"fmt"
)

var db *gorm.DB

func migrate_db() error {
	var m common.Maintenance
	db , err := gorm.Open("mysql","root:root@tcp(localhost:3306)/maint")
	if err != nil {
		return fmt.Errorf("Couldn't connect to the server")
	}
	ok := db.HasTable(&m)
	//fmt.Println(ok)
	if !ok {
		fmt.Println(ok)
		db.CreateTable(&m)
	}else {
		db.AutoMigrate(&m)
	}
	return nil
}
//func (common.Maint_DO) TableName() string {
//	return "maint"
//}
//func (common.Maintenance ) TableName() string {
//	return "Maintenance"
//}

func DBConnect() (err error)  {
	db , err = gorm.Open("mysql","root:root@tcp(localhost:3306)/maint?parseTime=true")
	if err != nil {
		return fmt.Errorf("Failed to connect to DB")
	}
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to DB")
	return nil
}

func GetDbConnection() (*gorm.DB, error) {
	//log.Println("Getting a DB connection")
	return db, nil
}