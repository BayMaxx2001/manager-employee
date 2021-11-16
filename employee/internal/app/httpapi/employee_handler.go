package httpapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/BayMaxx2001/manager-employee/employee/internal/service"
	"github.com/BayMaxx2001/manager-employee/pkg/httputil"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := service.GetAllEmployees(r.Context())
	if err != nil {
		log.Println("Error at GetAllEmployees of handlers/handlers.go", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		return
	}

	json.NewEncoder(w).Encode(employees)
	httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Success))
}

// Create new employee
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var payload service.AddEmployeeCommand
	if err := httputil.BindJsonReq(r, &payload); err != nil {
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		fmt.Println(err)
		return
	}

	employee, err := service.AddEmployee(r.Context(), payload)
	if err != nil {
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		return
	}

	json.NewEncoder(w).Encode(employee)
	httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Success))
}

// Delete Employee
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uid")

	//delete row in database
	err := service.DeleteEmployeeByUID(r.Context(), service.DeleteEmployeeByUIDCommand(id))
	if err != nil {
		log.Println("Error at DeleteEmployee of pkg/handlers/handlers.go when delete", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		return
	}

	httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Success))
}

// Find employee
func FindEmployee(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "uid")

	employee, err := service.FindEmployeeByUID(r.Context(), service.FindEmployeeByUIDCommand(idString))
	if err != nil {
		log.Println("Error at FindEmployee of pkg/handlers/handlers.go when search", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		return
	}

	json.NewEncoder(w).Encode(employee)
	httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Success))
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var payload service.UpdateEmployeeCommand
	if err := httputil.BindJsonReq(r, &payload); err != nil {
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		return
	}

	payload.UID = chi.URLParam(r, "uid")
	err := service.UpdateEmployeeById(r.Context(), payload)
	if err != nil {
		log.Println("Error at UpdateEmployee of pkg/handlers/handlers.go when update", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		return
	}

	json.NewEncoder(w).Encode(payload)
	httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Success))
}
