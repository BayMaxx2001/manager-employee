package service

import (
	"context"

	"github.com/BayMaxx2001/manager-employee/team/internal/model"
	"github.com/BayMaxx2001/manager-employee/team/internal/persistence"
	"github.com/asaskevich/govalidator"
)

type GetAllTeamCommand struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

func NewGetAllTeamCommand(e model.Team) GetAllTeamCommand {
	return GetAllTeamCommand{
		UID:  e.UID,
		Name: e.Name,
	}
}

func (c GetAllTeamCommand) Valid() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

func GetAllTeams(ctx context.Context) (teams []model.Team, err error) {
	teams, err = persistence.Teams().GetAll(ctx)
	return
}
