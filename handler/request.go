package handler

import "fmt"

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (cpr *CreateOpeningRequest) Validate() error {
	if cpr.Role == "" && cpr.Company == "" && cpr.Location == "" && cpr.Link == "" && cpr.Remote == nil && cpr.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if cpr.Role == "" {
		return errParamsIsRequired("role", "string")
	}
	if cpr.Company == "" {
		return errParamsIsRequired("company", "string")
	}
	if cpr.Location == "" {
		return errParamsIsRequired("location", "string")
	}
	if cpr.Link == "" {
		return errParamsIsRequired("link", "string")
	}
	if cpr.Remote == nil {
		return errParamsIsRequired("remote", "bool")
	}
	if cpr.Salary <= 0 {
		return errParamsIsRequired("salary", "int64")
	}
	return nil
}
