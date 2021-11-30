package httpapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BayMaxx2001/manager-employee/pkg/httputil"
	"github.com/BayMaxx2001/manager-employee/team/internal/service"
	"github.com/go-chi/chi"
)

func GetAllTeams(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	Teams, err := service.GetAllTeams(r.Context())
	if err != nil {
		log.Println("Error at GetAllEmployees of handlers/handlers.go", err)
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
		Data:    Teams,
	})
}

// Create new team
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var payload service.AddTeamCommand

	if err := httputil.BindJsonReq(r, &payload); err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)

		fmt.Println(err)
		return
	}

	team, err := service.AddTeam(r.Context(), payload)
	if err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
		Data:    team,
	})
}

// Delete Team
func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	id := chi.URLParam(r, "uid")
	//delete row in database
	err := service.DeleteTeamByUID(r.Context(), service.DeleteTeamByUIDCommand(id))
	if err != nil {
		log.Println("Error at DeleteTeam of pkg/handlers/handlers.go when delete", err)
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
	})
}

// Find team
func FindTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	idString := chi.URLParam(r, "uid")

	team, err := service.FindTeamByUID(r.Context(), service.FindTeamByUIDCommand(idString))
	if err != nil {
		log.Println("Error at FindTeam of pkg/handlers/handlers.go when search", err)
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
		Data:    team,
	})
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	var (
		payload service.UpdateTeamCommand
	)
	if err := httputil.BindJsonReq(r, &payload); err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	idString := chi.URLParam(r, "uid")
	payload.UID = idString

	err := service.UpdateTeamById(r.Context(), payload)
	if err != nil {
		log.Println("Error at UpdateTeam of pkg/handlers/handlers.go when update", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
		Data:    payload,
	})
}

func AddTeamToEmploye(w http.ResponseWriter, r *http.Request) {
	event := chi.URLParam(r, "event")
	fmt.Println(event)

	var payload service.TeamAddEmployeeCommand
	if err := httputil.BindJsonReq(r, &payload); err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		fmt.Println(err)
		return
	}

	fmt.Println(payload.EmployeeId, "----", payload.TeamId)

	team, err := service.AddTeamToEmployee(r.Context(), payload)
	if err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
		Data:    team,
	})

}

func DeleteTeamToEmploye(w http.ResponseWriter, r *http.Request) {
	event := chi.URLParam(r, "event")
	fmt.Println(event)

	var payload service.TeamDeleteEmployeeCommand
	if err := httputil.BindJsonReq(r, &payload); err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		fmt.Println(err)
		return
	}

	fmt.Println(payload.EmployeeId, "----", payload.TeamId)

	err := service.DeleteTeamToEmployee(r.Context(), payload)
	if err != nil {
		httputil.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	_ = httputil.WriteJsonOK(w, httputil.ResponseBody{
		Message: httputil.Message(http.StatusOK, httputil.Success),
	})
}
