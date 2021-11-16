package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/BayMaxx2001/manager-employee/pkg/httputil"
	"github.com/BayMaxx2001/manager-employee/team/internal/config"
	"github.com/BayMaxx2001/manager-employee/team/internal/model"
	redis "gopkg.in/redis.v5"
)

type redisTeamRepository struct {
	server   string
	password string
	db       int
	Redis    *redis.Client
}

func newRedisClient(server string, password string, db int) *redis.Client {
	result := redis.NewClient(&redis.Options{
		Addr:         server,
		Password:     password,
		DB:           db,
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
	})
	return result
}

func NewRedisProvider(server string, password string, db int) *redisTeamRepository {
	redis := newRedisClient(server, password, db)
	if redis == nil {
		log.Fatalln("Redis server connected unsuccessfully")
	}
	return &redisTeamRepository{server, password, db, redis}
}

func newRedisTeamRepository() (repo TeamRepository, err error) {
	repo = NewRedisProvider(config.Get().RedisServer, config.Get().RedisPassword, config.Get().RedisDB)

	return repo, nil
}

func (repo *redisTeamRepository) FindByUID(ctx context.Context, uid string) (model.Team, error) {
	var res model.Team
	val, err := repo.get(uid)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (repo *redisTeamRepository) Save(ctx context.Context, team model.Team) error {
	err := repo.set(team.UID, httputil.BindJSON(team), time.Hour)
	fmt.Println("save redis", team.UID)
	return err
}

func (repo *redisTeamRepository) Update(ctx context.Context, uid string, team model.Team) error {
	rdStatus := repo.set(uid, httputil.BindJSON(team), time.Hour)
	return rdStatus
}

func (repo *redisTeamRepository) Remove(ctx context.Context, uid string) error {
	err := repo.del(uid)
	return err
}

func (repo *redisTeamRepository) GetAll(ctx context.Context) (teams []model.Team, err error) {
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = repo.Redis.Scan(cursor, "*", 0).Result()
		if err != nil {
			panic(err)
		}
		for _, key := range keys {
			team, err := repo.FindByUID(ctx, key)
			if err != nil {
				continue
			}
			teams = append(teams, team)
		}
		if cursor == 0 { // no more keys
			break
		}
	}
	return
}

// redis
func (r *redisTeamRepository) set(key string, val interface{}, expire time.Duration) error {
	return r.Redis.Set(key, val, expire).Err()
}

func (r *redisTeamRepository) get(key string) (string, error) {
	return r.Redis.Get(key).Result()
}

func (r *redisTeamRepository) del(key ...string) error {
	return r.Redis.Del(key...).Err()
}
