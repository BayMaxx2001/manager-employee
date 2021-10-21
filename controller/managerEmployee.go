package controller

import (
	"errors"
	"fmt"
	"github/BayMaxx2001/manager-employees/model"
	"github/BayMaxx2001/manager-employees/scan"
)

type ListEmployee struct {
	LsEmployee []model.Employee
}

func (ls *ListEmployee) CreateEmployee() error {
	var employee model.Employee

	//set id of employee
	fmt.Print("Input ID: ")
	id, err := ls.setID(&employee)
	if err != nil {
		return err
	}

	// set name of employee
	fmt.Print("Input name: ")
	err = ls.SetNameById(id, &employee)
	if err != nil {
		return err
	}

	// set gender of employee
	fmt.Print("Input gender: ")
	err = ls.SetGenderById(id, &employee)
	if err != nil {
		return err
	}

	// set date of birth of employee
	fmt.Print("Input dob follow format yyyy-MM-dd: ")
	err = ls.SetDobById(id, &employee)
	if err != nil {
		return err
	}

	// append to list
	ls.LsEmployee = append(ls.LsEmployee, employee)
	return nil
}

func (ls *ListEmployee) GetAllEmployee() {
	fmt.Printf("Number of Employee: %d\n", len(ls.LsEmployee))
	fmt.Printf("|%-5s|%-5s|%-10s|%-10s|%-10s\n", "STT", "ID", "NAME", "GENDER", "DOB")
	for i, employee := range ls.LsEmployee {
		fmt.Printf("|%-5d", i+1)
		fmt.Printf("|%-5d|%-10s|%-10s|%-10s\n", employee.GetId(), employee.GetName(), employee.GetGender(), employee.GetDob().Format("02-Jan-2006"))
	}
}

func (ls *ListEmployee) UpdateByID(id int) error {
	employee, err := ls.GetEmployeeById(id)
	if err != nil {
		return err
	}
	fmt.Println("Information of employee")
	fmt.Printf("|%-5d|%-10s|%-10s|%-10s\n", employee.GetId(), employee.GetName(), employee.GetGender(), employee.GetDob().Format("02-Jan-2006"))
	fmt.Println("Update :")

	//update name of employee
	fmt.Print("Input name: ")
	err = ls.SetNameById(id, &employee)
	if err != nil {
		return err
	}

	// update gender of employee
	fmt.Print("Input gender: ")
	err = ls.SetGenderById(id, &employee)
	if err != nil {
		return err
	}

	// update date of birth of employee
	fmt.Print("Input dob follow format yyyy-MM-dd: ")
	err = ls.SetDobById(id, &employee)
	if err != nil {
		return err
	}
	return nil
}

func (ls *ListEmployee) DeleteByID(id int) error {
	if !scan.IsExist(&ls.LsEmployee, id) {
		err := errors.New("id not exist")
		return err
	}
	ls.LsEmployee = append(ls.LsEmployee[:id-1], ls.LsEmployee[id:]...)
	return nil
}

func (ls *ListEmployee) GetEmployeeById(id int) (model.Employee, error) {
	if len(ls.LsEmployee) == 0 {
		err := errors.New("list of employee is empty")
		return model.Employee{}, err
	}
	if !scan.IsExist(&ls.LsEmployee, id) {
		err := errors.New("id not exist")
		return model.Employee{}, err
	}

	for _, employee := range ls.LsEmployee {
		if employee.GetId() == id {
			return employee, nil
		}
	}
	return model.Employee{}, nil
}
