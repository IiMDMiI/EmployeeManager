package employeeManagement

import (
	aerr "github.com/IiMDMiI/employeeManager/internal/appErrors"
)

type Problem struct {
	Type    string
	Message string
	Details string
}

func NewInvalidBodyProblem(err error) Problem {
	return Problem{
		Type:    aerr.BadArg,
		Message: "Invalid request body",
		Details: err.Error()}
}

func NewInvalidQueryProblem(details string) Problem {
	return Problem{
		Type:    aerr.BadArg,
		Message: "Invalid query parameters",
		Details: details}
}

func NewUnexpectedProblem(errType string) Problem {
	return Problem{
		Type:    errType,
		Message: "Unexpected eror",
		Details: "Try to repeat your request later"}
}

func NewUnauthorizedProblem() Problem {
	return Problem{
		Type:    aerr.Security,
		Message: "Incorrect username and password",
		Details: "Ensure that the username and password included in the request are correct"}
}
