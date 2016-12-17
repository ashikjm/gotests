package maintenance

import (
	"fmt"
)

func DestroyAllVMs(id int) error {
	VMList, _ := GetVmList(id)
	for _,k := range VMList {
		err := DestroyVM(k)
		if err != nil {
			//ToDo : Return the instanceid of the VM facing issues.
			return fmt.Errorf("Aborting due to errors")
		}
	}
	return nil
}

func DestroyVM(id int) error {
	_, err := fmt.Println("Destroyed VM")
	if err != nil {
		return err
	}
	return nil
}

func MigrateVM(id int) error  {
	_, err := fmt.Println("Migrated VM")
	if err != nil {
		return err
	}
	return nil
}
