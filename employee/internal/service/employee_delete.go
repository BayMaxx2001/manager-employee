package service

import (
	"context"

	"github.com/BayMaxx2001/manager-employee/employee/internal/persistence"
	"github.com/google/uuid"
)

type DeleteEmployeeByUIDCommand string

func (c DeleteEmployeeByUIDCommand) Valid() error {
	_, err := uuid.Parse(string(c))
	return err
}

func DeleteEmployeeByUID(ctx context.Context, command DeleteEmployeeByUIDCommand) error {
	if err := command.Valid(); err != nil {
		return err
	}

	err := persistence.Employees().Remove(ctx, string(command))
	return err
}
