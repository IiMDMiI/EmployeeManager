package employeeManagement

const (
	UnfilledId = -1
)

type Employee struct {
	Id         int
	Name       string
	Surname    string
	Phone      string
	CompanyId  int
	Passport   Passport
	Department Department
}

type Passport struct {
	Type   string
	Number string
}

type Department struct {
	Name  string
	Phone string
}

func NewEmptyEmploee() *Employee {
	emp := Employee{}
	emp.Id = UnfilledId
	emp.CompanyId = UnfilledId
	return &emp
}
