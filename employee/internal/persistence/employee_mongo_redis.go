package persistence

import (
	"context"
	"fmt"

	"github.com/BayMaxx2001/manager-employee/employee/internal/model"
)

type employeeMongoRedis struct {
	employeeCache EmployeeRepository
	employeeMongo EmployeeRepository
}

func LoadEmployeesMongoRedisRepository() error {
	employeeCache, err := newRedisEmployeeRepository()
	if err != nil {
		return err
	}

	employeeMongo, err := newMongoEmployeeRepository()
	if err != nil {
		return err
	}

	employees = &employeeMongoRedis{employeeCache, employeeMongo}
	return nil
}

func (e *employeeMongoRedis) FindByUID(ctx context.Context, uid string) (empl model.Employee, err error) {
	empl, err = e.employeeCache.FindByUID(ctx, uid)
	if err == nil {
		return
	}

	empl, err = e.employeeMongo.FindByUID(ctx, uid)
	if err != nil {
		return
	}

	e.employeeCache.Save(ctx, empl)
	return
}

func (e *employeeMongoRedis) Save(ctx context.Context, employee model.Employee) error {
	err := e.employeeMongo.Save(ctx, employee)
	if err != nil {
		return err
	}
	return nil
}

func (e *employeeMongoRedis) Update(ctx context.Context, uid string, employee model.Employee) error {
	err := e.employeeCache.Update(ctx, uid, employee)
	if err != nil {
		fmt.Println("Update into employee_redis", err)
		return err
	}

	err = e.employeeMongo.Update(ctx, uid, employee)
	if err != nil {
		fmt.Println("Update into employee_mongo", err)
		return err
	}
	return nil
}

func (e *employeeMongoRedis) Remove(ctx context.Context, uid string) error {
	err := e.employeeCache.Remove(ctx, uid)
	if err != nil {
		return err
	}

	err = e.employeeMongo.Remove(ctx, uid)
	if err != nil {
		return err
	}
	return nil
}

func (e *employeeMongoRedis) GetAll(ctx context.Context) (ls []model.Employee, err error) {
	return e.employeeMongo.GetAll(ctx)
}
