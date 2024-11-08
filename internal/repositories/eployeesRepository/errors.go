package eployeesRepository

import (
	aerr "github.com/IiMDMiI/employeeManager/internal/appErrors"
	"github.com/IiMDMiI/employeeManager/internal/dbservice"
)

var (
	ErrBadPhoneNumberFormat = &aerr.AppError{
		Type:    aerr.BadArg,
		Message: "incorrect phone format",
	}
	ErrMissingEmployeeId = &aerr.AppError{
		Type:    aerr.BadArg,
		Message: "employee's Id is required",
	}
	ErrBadCompanyIdOrBadDepart = &aerr.AppError{
		Type:    aerr.BadArg,
		Message: "the company or department doesn't exist",
		Err:     dbservice.ErrForeignKeyViolation,
	}
	ErrEmployeeAlreadyRegistred = &aerr.AppError{
		Type:    aerr.BadArg,
		Message: "the employee is already registred",
		Err:     dbservice.ErrUniqueViolation,
	}
	ErrNoRecordsWereUpdated = &aerr.AppError{
		Type:    aerr.BadArg,
		Message: dbservice.ErrNoRecordsWereUpdated.Error(),
		Err:     dbservice.ErrNoRecordsWereUpdated,
	}
)
