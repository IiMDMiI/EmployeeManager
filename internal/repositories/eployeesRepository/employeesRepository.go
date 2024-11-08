package eployeesRepository

import (
	"errors"
	"log"

	em "github.com/IiMDMiI/employeeManager/api/employeeManagment"
	"github.com/IiMDMiI/employeeManager/internal/dbservice"
)

var valid = NewValidator()

type EmploeesRepositoryService interface {
	Create(emp *em.Employee) (int, error)
	Update(emp *em.Employee) error
	Delete(id int) error
	GetCompanyEmployees(companyId int) ([]em.Employee, error)
	GetDepartmentEmployees(companyId int, departmentId string) ([]em.Employee, error)
}

func New(db dbservice.DBService) EmploeesRepositoryService {
	return &emploeesRepository{db}
}

type emploeesRepository struct {
	db dbservice.DBService
}

func (er *emploeesRepository) Create(emp *em.Employee) (int, error) {
	if err := valid.Validate(emp); err != nil {
		return 0, err
	}

	id, err := er.db.CreateEmployee(emp)
	if err != nil {
		return id, er.dbErrorToAppError(err)
	}
	return id, nil
}

func (er *emploeesRepository) Update(emp *em.Employee) error {
	if emp.Id == em.UnfilledId {
		return ErrMissingEmployeeId
	}
	if emp.Phone != "" {
		if err := valid.ValidatePhone(emp.Phone); err != nil {
			return ErrBadPhoneNumberFormat
		}
	}

	err := er.db.UpdateEmployee(emp)
	if err != nil {
		return er.dbErrorToAppError(err)
	}
	return nil
}

func (er *emploeesRepository) Delete(id int) error {
	if err := er.db.DeleteEmployee(id); err != nil {
		return er.dbErrorToAppError(err)
	}
	return nil
}

func (er *emploeesRepository) GetCompanyEmployees(companyId int) ([]em.Employee, error) {
	employees, err := er.db.GetCompanyEmployees(companyId)
	if err != nil {
		return nil, er.dbErrorToAppError(err)
	}
	return employees, nil
}

func (er *emploeesRepository) GetDepartmentEmployees(companyId int, department string) ([]em.Employee, error) {
	employees, err := er.db.GetDepartmentEmployees(companyId, department)
	if err != nil {
		return nil, er.dbErrorToAppError(err)
	}
	return employees, nil
}

func (er *emploeesRepository) dbErrorToAppError(err error) error {
	if errors.Is(err, dbservice.ErrForeignKeyViolation) {
		return ErrBadCompanyIdOrBadDepart
	}
	if errors.Is(err, dbservice.ErrUniqueViolation) {
		return ErrEmployeeAlreadyRegistred
	}
	if errors.Is(err, dbservice.ErrNoRecordsWereUpdated) {
		return ErrNoRecordsWereUpdated
	}
	log.Println(err)
	return err
}
