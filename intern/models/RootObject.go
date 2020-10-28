package models

type Object struct {
	Key          string        `json:"root_key"`
	Values       []string      `json:"values"`
	ChildObjects []ChildObject `json:"child_objects"`
	TotalObjects int           `json:"total"`
}
