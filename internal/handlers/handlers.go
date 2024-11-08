package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	em "github.com/IiMDMiI/employeeManager/api/employeeManagment"
	aerr "github.com/IiMDMiI/employeeManager/internal/appErrors"
	mw "github.com/IiMDMiI/employeeManager/internal/middleware"
	er "github.com/IiMDMiI/employeeManager/internal/repositories/eployeesRepository"
)

var empsRep er.EmploeesRepositoryService

func SetUp(newUsersRep er.EmploeesRepositoryService) {
	empsRep = newUsersRep
	setRoutes()
}

func setRoutes() {
	prefix := "/api/v1"

	http.Handle("POST "+prefix+"/employee", mw.NewAuth(http.HandlerFunc(CreateEmployee)))
	http.Handle("DELETE "+prefix+"/employee", mw.NewAuth(http.HandlerFunc(DeleteEmployee)))
	http.Handle("PUT "+prefix+"/employee", mw.NewAuth(http.HandlerFunc(UpdateEmployee)))

	http.HandleFunc("GET "+prefix+"/employees", GetCompanyEmployees)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	emp := em.NewEmptyEmploee()

	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		setResponseWithInvalidBody(w, err)
		return
	}

	id, err := empsRep.Create(emp)
	if err != nil {
		var errValidation *er.ValidationError
		if errors.As(err, &errValidation) {
			setResponseWithInvalidBody(w, err)
		} else {
			handleCommonErrors(err, w)
		}
		return
	}

	setResponse(w, http.StatusOK, struct {
		Id string `json:"id"`
	}{Id: fmt.Sprint(id)})
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		setResponseWithProblem(w, http.StatusBadRequest,
			em.NewInvalidQueryProblem("Employee Id is required and must be an integer"))
		return
	}
	if err2 := empsRep.Delete(id); err2 != nil {
		handleCommonErrors(err2, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	emp := em.NewEmptyEmploee()
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		setResponseWithInvalidBody(w, err)
		return
	}

	if err := empsRep.Update(emp); err != nil {
		switch {
		case errors.Is(err, er.ErrMissingEmployeeId):
			setResponseWithInvalidBody(w, err)
		case errors.Is(err, er.ErrBadPhoneNumberFormat):
			setResponseWithInvalidBody(w, err)
		default:
			handleCommonErrors(err, w)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetCompanyEmployees(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		setResponseWithProblem(w, http.StatusBadRequest,
			em.NewInvalidQueryProblem("CompanyID must be an integer"))
		return
	}

	var emps []em.Employee
	var err2 error

	dep := r.URL.Query().Get("dep")
	if dep == "" {
		emps, err2 = empsRep.GetCompanyEmployees(id)
	} else {
		emps, err2 = empsRep.GetDepartmentEmployees(id, dep)
	}

	shouldReturn := handleGetEmpsDBErrors(err2, w, emps)
	if shouldReturn {
		return
	}

	setResponse(w, http.StatusOK, emps)
}

func handleGetEmpsDBErrors(err error, w http.ResponseWriter, emps []em.Employee) bool {
	if err != nil {
		setResponseWithProblem(w, http.StatusInternalServerError, em.NewUnexpectedProblem(aerr.InternalDB))
		return true
	}
	if len(emps) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return true
	}
	return false
}

func handleCommonErrors(err error, w http.ResponseWriter) {
	switch {
	case errors.Is(err, er.ErrBadCompanyIdOrBadDepart):
		setResponseWithInvalidBody(w, err)
	case errors.Is(err, er.ErrEmployeeAlreadyRegistred):
		setResponseWithInvalidBody(w, err)
	case errors.Is(err, er.ErrNoRecordsWereUpdated):
		setResponseWithInvalidBody(w, err)
	default:
		setResponseWithProblem(w, http.StatusInternalServerError, em.NewUnexpectedProblem(aerr.InternalDB))
	}
}

func setResponse(w http.ResponseWriter, status int, responseBody any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(responseBody)
}

func setResponseWithProblem(w http.ResponseWriter, status int, problem em.Problem) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(problem)
}

func setResponseWithInvalidBody(w http.ResponseWriter, err error) {
	setResponseWithProblem(w, http.StatusBadRequest, em.NewInvalidBodyProblem(err))
}
