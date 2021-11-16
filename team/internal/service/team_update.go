package service

import (
	"context"

	"github.com/BayMaxx2001/manager-employee/team/internal/model"
	"github.com/BayMaxx2001/manager-employee/team/internal/persistence"
	"github.com/asaskevich/govalidator"
)

type UpdateTeamCommand struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

func NewUpdateTeamCommand(e model.Team) UpdateTeamCommand {
	return UpdateTeamCommand{
		UID:  e.UID,
		Name: e.Name,
	}
}
func (c UpdateTeamCommand) Valid() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

func UpdateTeamById(ctx context.Context, command UpdateTeamCommand) error {
	if err := command.Valid(); err != nil {
		return err
	}

	team := model.Team{
		UID:  command.UID,
		Name: command.Name,
	}

	err := persistence.Teams().Update(ctx, command.UID, team)
	return err
}
