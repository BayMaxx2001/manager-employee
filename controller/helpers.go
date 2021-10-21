package controller

import (
	"errors"
	"github/BayMaxx2001/ManagerEmployees/model"
	"github/BayMaxx2001/ManagerEmployees/scan"
)

//set id of employee
func (ls *ListEmployee) setID(employee *model.Employee) (int, error) {
	idStr := scan.ScannerString()
	idNum, err := scan.ValidateId(idStr)
	if err != nil {
		return 0, err
	}
	if scan.IsExist(&ls.LsEmployee, idNum) {
		err := errors.New("id already exists")
		return 0, err
	}
	employee.SetId(idNum)
	return idNum, nil
}

// Set name of employee ID
func (ls *ListEmployee) SetNameById(id int, employee *model.Employee) error {
	name := scan.ScannerString()
	name, err := scan.ValidateName(name)
	if err != nil {
		err := errors.New("name is empty")
		return err
	}
	employee.SetName(name)
	return nil
}

// Set gender of employee ID
func (ls *ListEmployee) SetGenderById(id int, employee *model.Employee) error {
	gender := scan.ScannerString()
	gender, err := scan.ValidateGender(gender)
	if err != nil {
		err := errors.New("gender is empty")
		return err
	}
	employee.SetGender(gender)
	return nil
}

// Set date of birth of employee
func (ls *ListEmployee) SetDobById(id int, employee *model.Employee) error {
	dateS := scan.ScannerString()
	dateT, err := scan.ValidateDOB(dateS)
	if err != nil {
		return err
	}
	employee.SetDob(dateT)
	return nil
}
