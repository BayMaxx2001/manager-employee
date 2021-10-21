package main

import (
	"fmt"
	"github/BayMaxx2001/manager-employees/controller"
	"github/BayMaxx2001/manager-employees/menu"
	"github/BayMaxx2001/manager-employees/scan"
)

func main() {
	var (
		menu         menu.Menu
		ListEmployee controller.ListEmployee
		choice       int
	)
	menu.InitMenu()
	menu.DisplayMenu()
	for {
		fmt.Print("Enter choice: ")
		choice = scan.ScannerNumber()
		switch choice {
		case 1:
			err := ListEmployee.CreateEmployee()
			if err != nil {
				fmt.Println(err)
				continue
			}

		case 2:
			if scan.IsEmpty(&ListEmployee.LsEmployee) {
				continue
			}
			fmt.Print("Input ID: ")
			idNum := scan.ScannerNumber()
			err := ListEmployee.UpdateByID(idNum)
			if err != nil {
				fmt.Println(err)
				continue
			}

		case 3:
			if scan.IsEmpty(&ListEmployee.LsEmployee) {
				continue
			}
			fmt.Print("Input ID: ")
			idNum := scan.ScannerNumber()
			err := ListEmployee.DeleteByID(idNum)
			if err != nil {
				fmt.Println(err)
				continue
			}

		case 4:
			if scan.IsEmpty(&ListEmployee.LsEmployee) {
				continue
			}
			fmt.Print("Input ID: ")
			idNum := scan.ScannerNumber()
			employee, err := ListEmployee.GetEmployeeById(idNum)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf(" |%-5d|%-10s|%-10s|%-10s\n", employee.GetId(), employee.GetName(), employee.GetGender(), employee.GetDob().Format("02-Jan-2006"))
		case 5:
			if scan.IsEmpty(&ListEmployee.LsEmployee) {
				continue
			}
			ListEmployee.GetAllEmployee()
		default:
			fmt.Println("Exit")
			return
		}

	}
}
