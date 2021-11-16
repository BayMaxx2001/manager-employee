package service

import (
	"context"

	"github.com/BayMaxx2001/manager-employee/team/internal/model"
	"github.com/BayMaxx2001/manager-employee/team/internal/persistence"
	"github.com/google/uuid"
)

type FindTeamByUIDCommand string

func (c FindTeamByUIDCommand) Valid() error {
	_, err := uuid.Parse(string(c))
	return err
}

func FindTeamByUID(ctx context.Context, command FindTeamByUIDCommand) (team model.Team, err error) {
	if err = command.Valid(); err != nil {
		return
	}

	team, err = persistence.Teams().FindByUID(ctx, string(command))
	if err != nil {
		return
	}
	return
}
