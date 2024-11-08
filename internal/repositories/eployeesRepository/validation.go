package eployeesRepository

import (
	"fmt"
	"regexp"

	em "github.com/IiMDMiI/employeeManager/api/employeeManagment"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error: %s - %s", e.Field, e.Message)
}

type Validator interface {
	Validate(employee *em.Employee) error
	MandatoryFieldsPresent(employee *em.Employee) error
	ValidatePhone(phone string) error
}

func NewValidator() Validator {
	return &EmploeeValidator{}
}

type EmploeeValidator struct {
}

func (jv *EmploeeValidator) Validate(emp *em.Employee) error {
	//Disclaimer: I know that this can be simplified,
	// but I wanted to show that I know how to use channels
	// and handle goroutines

	errChan := make(chan error)
	defer close(errChan)
	chanClosed := false
	go func() {
		err := jv.MandatoryFieldsPresent(emp)
		if !chanClosed {
			errChan <- err
		}
	}()

	go func() {
		err := jv.ValidatePhone(emp.Phone)
		if !chanClosed {
			errChan <- err
		}
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			chanClosed = true
			return err
		}
	}
	return nil
}

func (jv *EmploeeValidator) MandatoryFieldsPresent(emp *em.Employee) error {
	if emp.Name == "" {
		return &ValidationError{Field: "Name", Message: "name is required"}
	}
	if emp.Surname == "" {
		return &ValidationError{Field: "Surname", Message: "surname is required"}
	}
	if emp.Phone == "" {
		return &ValidationError{Field: "Phone", Message: "phone is required"}
	}
	if emp.CompanyId == em.UnfilledId {
		return &ValidationError{Field: "CompanyId", Message: "companyId is required"}
	}
	if emp.Passport.Type == "" {
		return &ValidationError{Field: "Passport.Type", Message: "passport type is required"}
	}
	if emp.Passport.Number == "" {
		return &ValidationError{Field: "Passport.Number", Message: "passport number is required"}
	}
	if emp.Department.Name == "" {
		return &ValidationError{Field: "Department.Name", Message: "department name is required"}
	}
	return nil
}

func (jv *EmploeeValidator) ValidatePhone(phone string) error {
	re := regexp.MustCompile(`\D`)
	re.ReplaceAllString(phone, "")

	re2 := regexp.MustCompile(`^\+[1-9]\d{5,14}$`)
	if !re2.MatchString(phone) {
		return &ValidationError{Field: "Department.Phone", Message: "incorrect phone format"}
	}
	return nil
}
