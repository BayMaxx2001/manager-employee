package model

type Team struct {
	UID           string   `json:"uid"`
	Name          string   `json:"name"`
	ListEmployees []string `json:"listemployees"`
}
