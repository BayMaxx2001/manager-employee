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
	Teams, err := service.GetAllTeams(r.Context())
	if err != nil {
		log.Println("Error at GetAllEmployees of handlers/handlers.go", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		return
	}
	json.NewEncoder(w).Encode(Teams)
	httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Success))
}

// Create new team
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	var payload service.AddTeamCommand

	if err := httputil.BindJsonReq(r, &payload); err != nil {
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))

		fmt.Println(err)
		return
	}

	team, err := service.AddTeam(r.Context(), payload)
	if err != nil {
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		return
	}

	json.NewEncoder(w).Encode(team)
	httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Success))
}

// Delete Team
func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uid")
	//delete row in database
	err := service.DeleteTeamByUID(r.Context(), service.DeleteTeamByUIDCommand(id))
	if err != nil {
		log.Println("Error at DeleteTeam of pkg/handlers/handlers.go when delete", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		return
	}
	httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Success))
}

// Find team
func FindTeam(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "uid")

	team, err := service.FindTeamByUID(r.Context(), service.FindTeamByUIDCommand(idString))
	if err != nil {
		log.Println("Error at FindTeam of pkg/handlers/handlers.go when search", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		return
	}
	json.NewEncoder(w).Encode(team)
	httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Success))
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	var (
		payload service.UpdateTeamCommand
	)
	if err := httputil.BindJsonReq(r, &payload); err != nil {
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		return
	}

	idString := chi.URLParam(r, "uid")
	payload.UID = idString

	err := service.UpdateTeamById(r.Context(), payload)
	if err != nil {
		log.Println("Error at UpdateTeam of pkg/handlers/handlers.go when update", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Error))
		return
	}

	json.NewEncoder(w).Encode(payload)
	httputil.Respond(w, httputil.Message(http.StatusOK, httputil.Success))
}
