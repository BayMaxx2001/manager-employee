package service

import (
	"context"

	"github.com/BayMaxx2001/manager-employee/team/internal/persistence"
	"github.com/google/uuid"
)

type DeleteTeamByUIDCommand string

func (c DeleteTeamByUIDCommand) Valid() error {
	_, err := uuid.Parse(string(c))
	return err
}

func DeleteTeamByUID(ctx context.Context, command DeleteTeamByUIDCommand) error {
	if err := command.Valid(); err != nil {
		return err
	}

	err := persistence.Teams().Remove(ctx, string(command))
	return err
}
