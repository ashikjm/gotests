package maintenance

import (
	"fmt"
)


//func  (M *common.Maintenance) UpdateMothershipState(state string) error  {
func  UpdateMothershipState(id int, state string) error  {
	fmt.Println(id)
	switch state {
	case "reserved":
		fmt.Println("Reserving the Mothership")
		return nil
	case "active":
		fmt.Println("Activating the Mothership")
		return nil
	default:
		return fmt.Errorf("Invalid Mothership state")
	}
	return nil
}

func DestroyMothership(id int) error {
	_, err := fmt.Println("Destroyed Mothership")
	if err != nil {
		return err
	}
	return nil
}

func  CreateMothership(id int) error {
	_, err := fmt.Println("Created Mothership")
	if err != nil {
		return err
	}
	return nil
}

func  StopMothership(id int) error {
	_, err := fmt.Println("Stopped Mothership")
	if err != nil {
		return err
	}
	return nil
}

func  StartMothership(id int) error {
	_, err := fmt.Println("Started Mothership")
	if err != nil {
		return err
	}
	return nil
}

func GetVmList(id int) ([]int,error) {
	_, err := fmt.Println("Got VM list")
	if err != nil {
		return nil, err
	}
	return []int{1000,2000,3000}, nil
}