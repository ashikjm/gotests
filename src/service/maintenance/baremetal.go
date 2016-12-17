package maintenance

import (
	"fmt"
)

func DecommissionBM(id int) error {
	_, err := fmt.Println("Decommissioned Baremetal")
	if err != nil {
		return err
	}
	return nil
}

func InwardBM(id int) error {
	_, err := fmt.Println("Inwarded Baremetal")
	if err != nil {
		return err
	}
	return nil
}

