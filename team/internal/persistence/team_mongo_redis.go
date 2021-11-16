package persistence

import (
	"context"
	"fmt"

	"github.com/BayMaxx2001/manager-employee/team/internal/model"
)

type teamMongoRedis struct {
	teamCache TeamRepository
	teamMongo TeamRepository
}

func LoadTeamsRepository() error {
	teamCache, err := newRedisTeamRepository()
	if err != nil {
		return err
	}
	teamMongo, err := newMongoTeamRepository()
	if err != nil {
		return err
	}
	teams = &teamMongoRedis{teamCache, teamMongo}
	return nil
}

func (t *teamMongoRedis) FindByUID(ctx context.Context, uid string) (team model.Team, err error) {
	team, err = t.teamCache.FindByUID(ctx, uid)
	if err == nil {
		return
	}

	team, err = t.teamMongo.FindByUID(ctx, uid)
	if err != nil {
		return
	}
	t.teamCache.Save(ctx, team)
	return
}

func (t *teamMongoRedis) Save(ctx context.Context, team model.Team) error {
	err := t.teamMongo.Save(ctx, team)
	if err != nil {
		return err
	}
	return nil
}

func (t *teamMongoRedis) Update(ctx context.Context, uid string, team model.Team) error {
	err := t.teamCache.Update(ctx, uid, team)
	if err != nil {
		fmt.Println("Update into team_redis", err)
		return err
	}
	err = t.teamMongo.Update(ctx, uid, team)
	if err != nil {
		fmt.Println("Update into team_mongo", err)
		return err
	}
	return nil
}

func (t *teamMongoRedis) Remove(ctx context.Context, uid string) error {
	err := t.teamCache.Remove(ctx, uid)
	if err != nil {
		return err
	}
	err = t.teamMongo.Remove(ctx, uid)
	if err != nil {
		return err
	}
	return nil
}

func (t *teamMongoRedis) GetAll(ctx context.Context) (ls []model.Team, err error) {
	ls, err = t.teamMongo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return
}
