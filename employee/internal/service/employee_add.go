package service

import (
	"context"

	"github.com/BayMaxx2001/manager-employee/employee/internal/model"
	"github.com/BayMaxx2001/manager-employee/employee/internal/persistence"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type AddEmployeeCommand struct {
	Name   string `json:"name"`
	DOB    string `json:"dob"`
	Gender int    `json:"gender"`
}

func NewAddEmployeeCommand(e model.Employee) AddEmployeeCommand {
	return AddEmployeeCommand{
		Name:   e.Name,
		DOB:    e.DOB,
		Gender: e.Gender,
	}
}

func (c AddEmployeeCommand) Valid() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

func AddEmployee(ctx context.Context, command AddEmployeeCommand) (employee model.Employee, err error) {
	if err = command.Valid(); err != nil {
		return
	}

	employee = model.Employee{
		UID:       uuid.NewString(),
		Name:      command.Name,
		DOB:       command.DOB,
		Gender:    command.Gender,
		ListTeams: []string{},
	}

	err = persistence.Employees().Save(ctx, employee)
	return
}
