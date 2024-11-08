package app

import (
	"net/http"

	"github.com/IiMDMiI/employeeManager/internal/dbservice"
	"github.com/IiMDMiI/employeeManager/internal/handlers"
	er "github.com/IiMDMiI/employeeManager/internal/repositories/eployeesRepository"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	db := dbservice.New()
	defer db.Close()

	emploeesRepository := er.New(db)
	handlers.SetUp(emploeesRepository)

	port := ":8080"
	println("Server is running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
	return nil
}
