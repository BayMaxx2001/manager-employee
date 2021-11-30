package httpapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BayMaxx2001/manager-employee/employee/internal/service"
	"github.com/BayMaxx2001/manager-employee/pkg/httputil"
	"github.com/BayMaxx2001/manager-employee/pkg/messaging/httppub"
	"github.com/go-chi/chi"
)

// get all employee
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	employees, err := service.GetAllEmployees(r.Context())
	if err != nil {
		log.Println("Error at GetAllEmployees of handlers/handlers.go", err)
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
		Data:    employees,
	})
}

// Create new employee
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var payload service.AddEmployeeCommand
	if err := httputil.BindJsonReq(r, &payload); err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		fmt.Println(err)
		return
	}

	employee, err := service.AddEmployee(r.Context(), payload)
	if err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
		Data:    employee,
	})
}

// Delete Employee
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uid")

	//delete row in database
	err := service.DeleteEmployeeByUID(r.Context(), service.DeleteEmployeeByUIDCommand(id))
	if err != nil {
		log.Println("Error at DeleteEmployee of pkg/handlers/handlers.go when delete", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
	})
}

// Find employee
func FindEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	idString := chi.URLParam(r, "uid")

	employee, err := service.FindEmployeeByUID(r.Context(), service.FindEmployeeByUIDCommand(idString))
	if err != nil {
		log.Println("Error at FindEmployee of pkg/handlers/handlers.go when search", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
		Data:    employee,
	})
}

//update
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var payload service.UpdateEmployeeCommand
	if err := httputil.BindJsonReq(r, &payload); err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	payload.UID = chi.URLParam(r, "uid")
	err := service.UpdateEmployeeById(r.Context(), payload)
	if err != nil {
		log.Println("Error at UpdateEmployee of pkg/handlers/handlers.go when update", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
		Data:    payload,
	})
}

// add team
func AddEmployeeToTeam(w http.ResponseWriter, r *http.Request) {
	event := chi.URLParam(r, "event")
	fmt.Println(event)

	var payload service.EmployeeAddTeamCommand
	if err := httputil.BindJsonReq(r, &payload); err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	fmt.Println(payload.EmployeeId, "----", payload.TeamId)

	employee, err := service.AddEmployeeToTeam(r.Context(), payload)
	if err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
		Data:    employee,
	})
	httppub.Publish(&employee, http.MethodPost)
}

func DeleteEmployeeToTeam(w http.ResponseWriter, r *http.Request) {
	event := chi.URLParam(r, "event")
	fmt.Println(event)

	var payload service.EmployeeDeleteTeamCommand
	if err := httputil.BindJsonReq(r, &payload); err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		fmt.Println(err)
		return
	}

	fmt.Println(payload.EmployeeId, "----", payload.TeamId)

	err := service.DeleteEmployeeToTeam(r.Context(), payload)
	if err != nil {

		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
	})

	httppub.Publish(&payload, http.MethodDelete)
}
