package model

import "time"

const (
	GENDER_UNKNOWN = iota
	GENDER_MALE
	GENDER_FEMALE
)

type Employee struct {
	UID       string   `json:"uid"`
	Name      string   `json:"name"`
	DOB       string   `json:"dob"`
	Gender    int      `json:"gender"`
	ListTeams []string `json:"list_team"`
}

func (s *Employee) GenderStr() string {
	switch s.Gender {
	case GENDER_MALE:
		return "Male"
	case GENDER_FEMALE:
		return "Female"
	default:
		return "Unknown"
	}
}

func (s *Employee) DobFormat(layout string) (t time.Time) {
	t, err := time.Parse(layout, s.DOB)
	if err != nil {
		return
	}
	return
}
