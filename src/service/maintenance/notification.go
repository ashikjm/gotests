package maintenance

import (
	"fmt"
	"service/common"
	//"service/maintenance"
)


func NotifyVMOwners(M *common.Maintenance, msg_type string) error {
	vm_list , _ := GetVmList(M.Maintenance_type_id)
	fmt.Println(vm_list)
	fmt.Println(msg_type)
	switch msg_type {
	case "maint_schedule":
		_, err := fmt.Println("Informed VM owners of the Maintenance Schedule")
		if err != nil {
			return err
		}
		return nil
	case "maint_complete":
		_, err := fmt.Println("Informed VM Owners of Maintenance Completion")
		if err != nil {
			return err
		}
		return nil
	case "start_vms":
		_, err := fmt.Println("Informed VM owners of start the VMs")
		if err != nil {
			return err
		}
		return nil
	case "no_vm_operations":
		_, err := fmt.Println("Informed VM owners not to perform any Kloud-cli operations")
		if err != nil {
			return err
		}
		return nil
	case "migrate_vms":
		_, err := fmt.Println("Informed VM owners to migrate the VMs")
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}