package scan

import (
	"errors"
	"fmt"
	"github/BayMaxx2001/manager-employees/model"
	"strconv"
	"strings"
	"time"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func standardizeString(str string) string {
	str = strings.TrimSpace(str)
	str = standardizeSpaces(str)
	return str
}

func IsExist(lsE *[]model.Employee, id int) bool {
	for _, employee := range *lsE {
		if id == employee.GetId() {
			return true
		}
	}
	return false
}

func IsEmpty(lsE *[]model.Employee) bool {
	if len(*lsE) == 0 {
		fmt.Println("list of employee is empty")
		return true
	}
	return false
}
func ValidateName(input string) (string, error) {
	if input == "" {
		return "", errors.New("name is empty")
	}
	input = standardizeString(input)
	input = strings.ToLower(input)
	input = strings.Title(input)
	return input, nil
}

func ValidateGender(input string) (string, error) {
	if input == "" {
		return "", errors.New("gender is empty")
	}

	input = standardizeString(input)
	input = strings.ToLower(input)
	if input != "male" && input != "female" {
		return "both", nil
	}
	return input, nil
}

func ValidateDOB(input string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func ValidateId(s string) (int, error) {
	number, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(number), nil
}
