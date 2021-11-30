package persistence

import (
	"context"

	"github.com/BayMaxx2001/manager-employee/team/internal/model"
)

var (
	teams TeamRepository
)

type TeamRepository interface {
	FindByUID(ctx context.Context, uid string) (model.Team, error)
	Save(ctx context.Context, team model.Team) error
	Update(ctx context.Context, uid string, Team model.Team) error
	Remove(ctx context.Context, uid string) error
	GetAll(ctx context.Context) ([]model.Team, error)
	AddTeamToEmplopyee(ctx context.Context, team model.Team, eid string) error
	DeleteTeamToEmployee(ctx context.Context, team model.Team, eid string) error
}

func Teams() TeamRepository {
	if teams == nil {
		panic("persistence: Team not initiated")
	}
	return teams
}
