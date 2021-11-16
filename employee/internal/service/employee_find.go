package service

import (
	"context"

	"github.com/BayMaxx2001/manager-employee/employee/internal/model"
	"github.com/BayMaxx2001/manager-employee/employee/internal/persistence"
	"github.com/google/uuid"
)

type FindEmployeeByUIDCommand string

func (c FindEmployeeByUIDCommand) Valid() error {
	_, err := uuid.Parse(string(c))

	return err
}

func FindEmployeeByUID(ctx context.Context, command FindEmployeeByUIDCommand) (Employee model.Employee, err error) {
	if err = command.Valid(); err != nil {
		return
	}

	employee, err := persistence.Employees().FindByUID(ctx, string(command))
	if err != nil {
		return employee, err
	}

	return employee, nil
}
