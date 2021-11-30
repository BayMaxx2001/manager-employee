package service

import (
	"context"

	"github.com/BayMaxx2001/manager-employee/team/internal/model"
	"github.com/BayMaxx2001/manager-employee/team/internal/persistence"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type AddTeamCommand struct {
	Name string `json:"name"`
}

func NewAddTeamCommand(t model.Team) AddTeamCommand {
	return AddTeamCommand{
		Name: t.Name,
	}
}

func (c AddTeamCommand) Valid() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

func AddTeam(ctx context.Context, command AddTeamCommand) (team model.Team, err error) {
	if err = command.Valid(); err != nil {
		return
	}
	team = model.Team{
		UID:           uuid.NewString(),
		Name:          command.Name,
		ListEmployees: []string{},
	}
	err = persistence.Teams().Save(ctx, team)
	return
}
