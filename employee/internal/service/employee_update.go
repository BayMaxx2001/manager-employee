package service

import (
	"context"

	"github.com/BayMaxx2001/manager-employee/employee/internal/model"
	"github.com/BayMaxx2001/manager-employee/employee/internal/persistence"
	"github.com/asaskevich/govalidator"
)

type UpdateEmployeeCommand struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	DOB    string `json:"dob"`
	Gender int    `json:"gender"`
}

func NewUpdateEmployeeCommand(e model.Employee) UpdateEmployeeCommand {
	return UpdateEmployeeCommand{
		UID:    e.UID,
		Name:   e.Name,
		DOB:    e.DOB,
		Gender: e.Gender,
	}
}
func (c UpdateEmployeeCommand) Valid() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

func UpdateEmployeeById(ctx context.Context, command UpdateEmployeeCommand) error {
	if err := command.Valid(); err != nil {
		return err
	}

	employee := model.Employee{
		UID:    command.UID,
		Name:   command.Name,
		DOB:    command.DOB,
		Gender: command.Gender,
	}

	err := persistence.Employees().Update(ctx, command.UID, employee)
	return err
}
