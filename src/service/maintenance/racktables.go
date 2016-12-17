package maintenance

import (
	"fmt"
	//"service/maintenance"
	//"service/common"
)


//func (M *maintenance.Maint_obj) UpdateRacktables(action string, msg []string) error {
func UpdateRacktables(id int, action string, msg []string) error {
	err := UpdateStatusRacktables(action)
	if err != nil {
		return err
	}
	err = AddRacktablesMessage(msg)
	if err != nil {
		ToggleStatusRacktables()
		return err
	}
	return nil
}

func UpdateStatusRacktables(action string) error  {
	//ToDo: Below line will be replaced by actions to update the status
	_, err := fmt.Println("Funtion to perform %s added successfully" ,action)
	if err != nil {
		return fmt.Errorf("Couldn't Update")
	}
	return nil
}

func AddRacktablesMessage(msg []string) error {
	_, err := fmt.Println("Messages added to racktables successfully %v", msg)
	if err != nil {
		return fmt.Errorf("Couldn't Add the message")
	}
	return nil
}

func ToggleStatusRacktables() {
	//Below line will be replaced by actions to toggle the status
	fmt.Printf("Function to toggle")
}