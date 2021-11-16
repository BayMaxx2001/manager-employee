package service

import (
	"context"

	"github.com/BayMaxx2001/manager-employee/employee/internal/model"
	"github.com/BayMaxx2001/manager-employee/employee/internal/persistence"
	"github.com/asaskevich/govalidator"
)

type GetAllEmployeesCommand struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	DOB    string `json:"dob"`
	Gender int    `json:"gender"`
}

func NewGetAllEmployeesCommand(e model.Employee) GetAllEmployeesCommand {
	return GetAllEmployeesCommand{
		UID:    e.UID,
		Name:   e.Name,
		DOB:    e.DOB,
		Gender: e.Gender,
	}
}

func (c GetAllEmployeesCommand) Valid() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

func GetAllEmployees(ctx context.Context) (employees []model.Employee, err error) {
	employees, err = persistence.Employees().GetAll(ctx)
	return
}
