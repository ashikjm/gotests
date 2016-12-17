package base

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"service/common"
	"fmt"
)

//type Maint struct {
//	id int `json:"id"`
//	resource_type string `json:"resource_type"`
//}

func GetAnObject(id string) (*common.Maintenance, error)  {
	var M common.Maintenance
	db, err := gorm.Open("mysql","root:root@tcp(localhost:3306)/maint?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("Couldn't connect to DB")
	}
	defer db.Close()
	fmt.Println("%v", id)
	//db.Where("id = ?", id).Find(&M)
	db.Find(&M, id)
	//for _ , k := range(M) {
	//	fmt.Printf("%T", k)
	//	fmt.Printf("%v", k)
	//}
	return &M, nil
}

func UpdateObjectValue(id int16, value string) error {
	//ToDo: Transaction Not Working
	fmt.Println("id is %v",id)
	var M common.Maintenance
	M.ID = id
	//M.Status = "scheduled"
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/maint?parseTime=true")
	if err != nil {
		return err
	}
	defer db.Close()
	db.Find(&M)
	M.Status = value
	fmt.Println(M.CreatedAt)
	//tx := db.Begin()
	//if err := tx.Model(&M).Update("status", value); err != nil {
	//if err := tx.Model(&maint_obj).Update("status", value) ; err != nil {
	//	fmt.Println("Printing Error")
	//	fmt.Println(err)
	//	tx.Rollback()
	//	return fmt.Errorf("Couldn't Update. Rolling Back")
	//}
	//tx.Commit()
	db.Save(&M)
	//db.Model(&M).Where("ID = ?", id).Update("status", value)
	fmt.Println(M)
	return nil
}

func CreateMaintenanceObject(M common.Maintenance) (int16, error)  {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/maint?parseTime=true")
	if err != nil {
		return 0 , fmt.Errorf("Couldn't connect to mysql")
	}
	//db, err := GetDbConnection()
	//if err != nil {
	//	fmt.Println("Couldn't Open")
	//}
	defer db.Close()
	//fmt.Printf("%#v",M)
	//db.Create(&M)
	tx := db.Begin()
	if err := tx.Create(&M).Error; err != nil {
		tx.Rollback()
		return 0 , err
	}
	tx.Commit()
	//ok := db.NewRecord(M)
	//if !ok {
	//	fmt.Errorf("Couldn't update the record")
	//}
	return M.ID, nil
}


//func GetAllObjects() (*[]common.Maint_DO) {
//	var m_n []common.Maint_DO
//	db, err := gorm.Open("mysql","root:root@tcp(localhost:3306)/maint")
//	if err != nil {
//		return nil
//	}
//	//db, err := GetDbConnection()
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	fmt.Printf("Opened the conn")
//	defer db.Close()
//	fmt.Println("Finding")
//	//db.Create(&Maint_DO{})
//	db.Find(&m_n)
//	fmt.Println(m_n)
//	fmt.Printf("**Found")
//	return &m_n
//}