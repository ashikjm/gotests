package base

import (
	"net/http"
	//"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"service/maintenance"
	"github.com/gorilla/mux"
	//"service/common"
	"service/common"
	"github.com/jinzhu/gorm"
)


//func GetAllMaintenanceHandlerTest(r *http.Request) ([]byte, *appError) {
//	m_o := GetAllObjects()
//	fmt.Printf("Marshalling")
//	a, _ := json.Marshal(&m_o)
//	//a := []byte("Success accessing this")
//	//(*writer).Write(a)
//	return a, nil
//}

//func CreateMaintenanceHandlerTest(r *http.Request) ([]byte, *appError) {
//	var decoder_obj common.Maint_DO
//	err := json.NewDecoder(r.Body).Decode(&decoder_obj)
//	if err != nil {
//		return []byte{}, &appError{400, err.Error()}
//	}
//	err = CreateMaintenanceObject(decoder_obj)
//	fmt.Println(err)
//	if err != nil {
//		return []byte("There was an error"), &appError{500, err.Error()}
//	}
//	return []byte("Created an entry"), &appError{200, err.Error()}
//}

//func GetMaintenanceHandlerTest(r *http.Request) ([]byte, *appError) {
//	vars := mux.Vars(r)
//	m_o := GetAnObject(vars["id"])
//	a, _ := json.Marshal(&m_o)
//	return a , nil
//}

func GetMaintenanceHandler(r *http.Request) ([]byte, *appError) {
	vars := mux.Vars(r)
	m_o, err := GetAnObject(vars["id"])
	if err != nil {
		return []byte{}, &appError{400, err.Error()}
	}
	a, _ := json.Marshal(&m_o)
	return a , nil
}

func CreateMaintenanceHandler(r *http.Request) ([]byte, *appError) {
	var decoder_obj common.Maintenance
	err := json.NewDecoder(r.Body).Decode(&decoder_obj)
	if err != nil {
		return []byte{}, &appError{400, err.Error()}
	}
	fmt.Println(decoder_obj)
	id , err := CreateMaintenanceObject(decoder_obj)
	//fmt.Println(err)
	if err != nil {
		return []byte("There was an error"), &appError{500, err.Error()}
	}
	fmt.Println(id)
	return []byte(`{"message": "Maintenance Object Created"}`), nil
}

func ScheduleMaintenanceHandler(r *http.Request) ([]byte, *appError) {
	vars := mux.Vars(r)
	maint_obj_id, _ := GetAnObject(vars["id"])
	//ToDo : Add conditional scheduling
	id := maint_obj_id.Resource_id
	switch maint_obj_id.Maintenance_type_id {
	case 1,2,3,4,6:
		//ToDo : Add Error Handling here
		maintenance.UpdateMothershipState(id, "reserved")
		msg :=  []string{"This is a message"}
		maintenance.UpdateRacktables(id, "add_issue", msg )
		maintenance.NotifyVMOwners(maint_obj_id, "maint_schedule")
		UpdateObjectValue(maint_obj_id.ID, "scheduled")
	case 5:
		maintenance.UpdateMothershipState(id, "reserved")
		maintenance.NotifyVMOwners(maint_obj_id, "no_vm_operations")
		UpdateObjectValue(maint_obj_id.ID, "scheduled")
	}
	return []byte(`{"message": "Scheduled the Maintenance"}`), nil

}

func StartMaintenanceHandler(r *http.Request) ([]byte , *appError)  {
	vars := mux.Vars(r)
	maint_obj_id, _ := GetAnObject(vars["id"])
	if maint_obj_id.Status != "scheduled" {
		return []byte(`{"message": "This state change not allowed"}`), nil
	}
	id := maint_obj_id.Resource_id
	switch maint_obj_id.Maintenance_type_id {
	case 1:
		maintenance.DestroyAllVMs(maint_obj_id.Resource_id)
		maintenance.DestroyMothership(id)
		maintenance.DecommissionBM(id)
		UpdateObjectValue(maint_obj_id.ID, "started")
	case 2:
		maintenance.DestroyAllVMs(maint_obj_id.Resource_id)
		maintenance.DestroyMothership(id)
		UpdateObjectValue(maint_obj_id.ID, "started")
	case 3:
		maintenance.StopMothership(id)
		UpdateObjectValue(maint_obj_id.ID, "started")
	case 4,5,6:
		UpdateObjectValue(maint_obj_id.ID, "started")
	}
	return []byte(`{"message": "Started the Maintenance"}`), nil
}

func CompleteMaintenanceHandler(r *http.Request)  ([]byte , *appError)  {
	vars := mux.Vars(r)
	maint_obj_id, _ := GetAnObject(vars["id"])
	if maint_obj_id.Status != "started" {
		return []byte(`{"message": "This state change not allowed"}`), nil
	}
	id := maint_obj_id.Resource_id
	switch maint_obj_id.Maintenance_type_id {
	case 1:
		msg :=  []string{"This is a message"}
		maintenance.UpdateRacktables(id, "resolve_issue", msg)
		maintenance.InwardBM(id)
		UpdateObjectValue(maint_obj_id.ID, "completed")
	case 2:
		msg :=  []string{"This is a message"}
		maintenance.UpdateRacktables(id, "resolve_issue", msg)
		maintenance.CreateMothership(id)
		UpdateObjectValue(maint_obj_id.ID, "completed")
	case 3:
		msg :=  []string{"This is a message"}
		maintenance.UpdateRacktables(id, "resolve_issue", msg)
		maintenance.StartMothership(id)
		maintenance.UpdateMothershipState(id, "active")
		maintenance.NotifyVMOwners(maint_obj_id,"start_vms")
		UpdateObjectValue(maint_obj_id.ID, "completed")
	case 4:
		msg :=  []string{"This is a message"}
		maintenance.UpdateRacktables(id, "resolve_issue", msg)
		maintenance.UpdateMothershipState(id,"active")
		maintenance.NotifyVMOwners(maint_obj_id,"maint_complete")
		UpdateObjectValue(maint_obj_id.ID, "completed")
	case 5:
		maintenance.UpdateMothershipState(id,"active")
		maintenance.NotifyVMOwners(maint_obj_id,"maint_complete")
		UpdateObjectValue(maint_obj_id.ID, "completed")
	case 6:
		msg :=  []string{"This is a message"}
		maintenance.UpdateRacktables(id, "resolve_issue", msg)
		maintenance.UpdateMothershipState(id,"active")
		maintenance.NotifyVMOwners(maint_obj_id,"maint_complete")
		UpdateObjectValue(maint_obj_id.ID, "completed")

	}
	return []byte(`{"message": "Maintenance Complete"}`), nil
}