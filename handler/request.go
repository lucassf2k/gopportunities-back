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

func (cor *CreateOpeningRequest) Validate() error {
	if cor.Role == "" && cor.Company == "" && cor.Location == "" && cor.Link == "" && cor.Remote == nil && cor.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if cor.Role == "" {
		return errParamsIsRequired("role", "string")
	}
	if cor.Company == "" {
		return errParamsIsRequired("company", "string")
	}
	if cor.Location == "" {
		return errParamsIsRequired("location", "string")
	}
	if cor.Link == "" {
		return errParamsIsRequired("link", "string")
	}
	if cor.Remote == nil {
		return errParamsIsRequired("remote", "bool")
	}
	if cor.Salary <= 0 {
		return errParamsIsRequired("salary", "int64")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (uor *UpdateOpeningRequest) Validate() error {
	if uor.Role != "" || uor.Company != "" || uor.Link != "" || uor.Location != "" || uor.Remote != nil || uor.Salary > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
