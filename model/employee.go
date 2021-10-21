package model

import "time"

type Employee struct {
	id     int
	name   string
	gender string
	dob    time.Time
}

func (e *Employee) NewEmployee(id int, name string, gender string, dob time.Time) {
	e.id = id
	e.name = name
	e.gender = gender
	e.dob = dob
}
func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) GetGender() string {
	return e.gender
}

func (e *Employee) GetDob() time.Time {
	return e.dob
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) SetGender(gerder string) {
	e.gender = gerder
}

func (e *Employee) SetDob(dob time.Time) {
	e.dob = dob
}
