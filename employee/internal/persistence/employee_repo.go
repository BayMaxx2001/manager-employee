package persistence

import (
	"context"

	"github.com/BayMaxx2001/manager-employee/employee/internal/model"
)

var (
	employees EmployeeRepository
)

type EmployeeRepository interface {
	FindByUID(ctx context.Context, uid string) (model.Employee, error)
	Save(ctx context.Context, employee model.Employee) error
	Update(ctx context.Context, uid string, employee model.Employee) error
	Remove(ctx context.Context, uid string) error
	GetAll(ctx context.Context) ([]model.Employee, error)
	AddEmployeeToTeam(ctx context.Context, employee model.Employee, tid string) error
	DeleteEmployeeToTeam(ctx context.Context, employee model.Employee, tid string) error
}

func Employees() EmployeeRepository {
	if employees == nil {
		panic("persistence: Employee not initiated")
	}
	return employees
}
