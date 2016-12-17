package common

import (
	//"github.com/jinzhu/gorm"
	"time"
)


//import "database/sql"

//type Maint struct {
//	id int `json:"id"`
//	resource_type string `json:"resource_type"`
//}

type Maint_DO struct {
	Id int
	Resource_type string
}

type Maintenance struct  {
	//gorm.Model
	ID        int16 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Resource_type string
	Resource_id int
	Maintenance_type_id int
	Status string
	Action string
	Parent_id int
}