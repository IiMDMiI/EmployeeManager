package dbservice

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	em "github.com/IiMDMiI/employeeManager/api/employeeManagment"

	"github.com/lib/pq"
)

func init() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

const (
	FOREIGN_KEY_VIOLATION = "23503"
	UNIQUE_VIOLATION      = "23505"
)

var (
	psqlInfo                string
	ErrForeignKeyViolation  = fmt.Errorf("foreign key violation")
	ErrUniqueViolation      = fmt.Errorf("unique violation")
	ErrNoRecordsWereUpdated = fmt.Errorf("no records were updated")
)

type DBService interface {
	CreateEmployee(emp *em.Employee) (int, error)
	DeleteEmployee(id int) error
	UpdateEmployee(emp *em.Employee) error
	GetCompanyEmployees(companyId int) ([]em.Employee, error)
	GetDepartmentEmployees(companyId int, departmentId string) ([]em.Employee, error)
	Close()
}

func New() DBService {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	return &DB{db}
}

type DB struct {
	db *sql.DB
}

func (db *DB) Close() {
	db.db.Close()
}

func (db *DB) CreateEmployee(emp *em.Employee) (int, error) {
	row := db.db.QueryRow(`INSERT INTO employee (name, surname, phone, company_id, department_name, pass_type, pass_number)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`, emp.Name, emp.Surname, emp.Phone, emp.CompanyId,
		emp.Department.Name, emp.Passport.Type, emp.Passport.Number)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return id, db.clarifyDBEror(err)
	}

	return id, nil
}

func (db *DB) DeleteEmployee(id int) error {
	return db.execWithUpdateCheck(fmt.Sprintf("DELETE FROM employee WHERE id = %d", id))
}

func (db *DB) UpdateEmployee(emp *em.Employee) error {
	empArgs := createEmploeeArgs(emp)
	if len(empArgs) > 0 {
		empQuery := createUpdateQuery("employee", empArgs, fmt.Sprintf("id = %d", emp.Id))
		return db.execWithUpdateCheck(empQuery)
	}

	return nil
}

func (db *DB) execWithUpdateCheck(empQuery string) error {
	result, err := db.db.Exec(empQuery)
	if err != nil {
		return db.clarifyDBEror(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return ErrNoRecordsWereUpdated
	}
	return nil
}

func (db *DB) GetCompanyEmployees(companyId int) ([]em.Employee, error) {
	rows, err := db.db.Query(`SELECT e.id, e.name, e.surname, e.phone, e.company_id, e.department_name,
		d.phone AS dep_phone, pass_type, pass_number
		FROM employee e
		JOIN department d ON e.department_name = d.name AND e.company_id = d.company_id
		WHERE e.company_id = $1;`, companyId)

	if err != nil {
		return nil, err
	}

	return db.rowsToEmps(rows)
}

func (db *DB) GetDepartmentEmployees(companyId int, department string) ([]em.Employee, error) {
	rows, err := db.db.Query(`SELECT e.id, e.name, e.surname, e.phone, e.company_id, e.department_name,
	    d.phone AS dep_phone, pass_type, pass_number
		FROM employee e
		JOIN department d ON e.department_name = d.name AND e.company_id = d.company_id
		WHERE e.company_id = $1 and e.department_name = $2;`, companyId, department)

	if err != nil {
		return nil, err
	}

	return db.rowsToEmps(rows)
}

func (db *DB) rowsToEmps(rows *sql.Rows) ([]em.Employee, error) {
	defer rows.Close()

	var emps []em.Employee
	for rows.Next() {
		var emp em.Employee
		if err := rows.Scan(&emp.Id, &emp.Name, &emp.Surname, &emp.Phone,
			&emp.CompanyId, &emp.Department.Name, &emp.Department.Phone,
			&emp.Passport.Type, &emp.Passport.Number); err != nil {
			return nil, err
		}
		emps = append(emps, emp)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return emps, nil
}

func (db *DB) clarifyDBEror(err error) error {
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == FOREIGN_KEY_VIOLATION {
			return ErrForeignKeyViolation
		}
		if pqErr.Code == UNIQUE_VIOLATION {
			return ErrUniqueViolation
		}
	}
	return err
}
