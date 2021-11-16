package persistence

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"gopkg.in/redis.v5"

	"github.com/BayMaxx2001/manager-employee/employee/internal/config"
	"github.com/BayMaxx2001/manager-employee/employee/internal/model"
	"github.com/BayMaxx2001/manager-employee/pkg/httputil"
)

type redisEmployeeRepository struct {
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

func newRedisEmployeeRepository() (repo EmployeeRepository, err error) {
	repo = NewRedisProvider(config.Get().RedisServer, config.Get().RedisPassword, config.Get().RedisDB)

	return repo, nil
}

func NewRedisProvider(server string, password string, db int) *redisEmployeeRepository {
	redis := newRedisClient(server, password, db)
	if redis == nil {
		log.Fatalln("Redis server connected unsuccessfully")
	}

	return &redisEmployeeRepository{server, password, db, redis}
}

func (repo *redisEmployeeRepository) FindByUID(ctx context.Context, uid string) (model.Employee, error) {
	var res model.Employee

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

func (repo *redisEmployeeRepository) Save(ctx context.Context, employee model.Employee) error {
	err := repo.set(employee.UID, httputil.BindJSON(employee), time.Hour)

	return err
}

func (repo *redisEmployeeRepository) Update(ctx context.Context, uid string, employee model.Employee) error {
	rdStatus := repo.set(uid, httputil.BindJSON(employee), time.Hour)
	return rdStatus
}

func (repo *redisEmployeeRepository) Remove(ctx context.Context, uid string) error {
	err := repo.del(uid)
	return err
}

func (repo *redisEmployeeRepository) GetAll(ctx context.Context) (ls []model.Employee, err error) {
	var cursor uint64
	for {
		var keys []string
		var err error

		keys, cursor, err = repo.Redis.Scan(cursor, "*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			empl, err := repo.FindByUID(ctx, key)
			if err != nil {
				continue
			}
			ls = append(ls, empl)
		}

		if cursor == 0 { // no more keys
			break
		}
	}
	return
}

// redis
func (r *redisEmployeeRepository) set(key string, val interface{}, expire time.Duration) error {
	return r.Redis.Set(key, val, expire).Err()
}

func (r *redisEmployeeRepository) get(key string) (string, error) {
	return r.Redis.Get(key).Result()
}

func (r *redisEmployeeRepository) del(key ...string) error {
	return r.Redis.Del(key...).Err()
}
