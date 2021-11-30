package httpapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/BayMaxx2001/manager-employee/pkg/messaging/httppub"
	"github.com/BayMaxx2001/manager-employee/pkg/messaging/httpsub"
	"github.com/BayMaxx2001/manager-employee/team/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(mux *chi.Mux) {
	mux.Route("/api/v1", func(r chi.Router) {
		mux.Use(middleware.Timeout(30 * time.Second))
		mux.Use(middleware.Recoverer)

		r.Route("/team", func(r chi.Router) {
			r.Get("/", GetAllTeams)
			r.Post("/", CreateTeam)
			r.Put("/{uid}", UpdateTeam)
			r.Delete("/{uid}", DeleteTeam)
			r.Get("/{uid}", FindTeam)
		})
		r.Route("/event", func(r chi.Router) {
			r.Post("/{event}", AddTeamToEmploye)
			r.Delete("/{event}", DeleteTeamToEmploye)
		})
	})
}

func Serve(ctx context.Context, addr string) (err error) {
	defer func() {
		log.Println("HTTP server stopped", err)
	}()

	r := chi.NewRouter()

	Routes(r)

	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
	}()

	srv := http.Server{
		Addr:    addr,
		Handler: r,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	errChan := make(chan error, 1)

	// pubsub
	httppub.Init()
	httpsub.Init()

	err = subscriberAdd(ctx)
	if err != nil {
		return
	}

	err = subscriberDelete(ctx)
	if err != nil {
		return
	}

	go func(ctx context.Context, errChan chan error) {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}(ctx, errChan)

	log.Printf("HTTP server started at %s\n", addr)

	select {
	case <-ctx.Done():
		return nil
	case err = <-errChan:
		return err
	}
}

func subscriberAdd(ctx context.Context) (err error) { // oke
	sub := httpsub.NewSubscriber("employee-team")
	httpsub.ConnectSub(*sub, "employee-team")
	go func() {
		for {
			data := <-sub.C
			var teamAddEmployeeCommand service.TeamAddEmployeeCommand
			err = json.Unmarshal(data, &teamAddEmployeeCommand)
			if err != nil {
				return
			}
			fmt.Println(teamAddEmployeeCommand.EmployeeId, "--", teamAddEmployeeCommand.TeamId, "---", err)
			service.AddTeamToEmployee(ctx, teamAddEmployeeCommand)
		}
	}()
	return nil
}

func subscriberDelete(ctx context.Context) (err error) { // oke
	sub := httpsub.NewSubscriber("employee-team")
	httpsub.ConnectSub(*sub, "employee-team")
	go func() {
		for {
			var teamDeleteEmployeeCommand service.TeamDeleteEmployeeCommand
			data := <-sub.C
			err = json.Unmarshal(data, &teamDeleteEmployeeCommand)
			if err != nil {
				return
			}
			fmt.Println(teamDeleteEmployeeCommand.EmployeeId, "--", teamDeleteEmployeeCommand.TeamId, "---", err)
			service.DeleteTeamToEmployee(ctx, teamDeleteEmployeeCommand)
		}
	}()
	return nil
}
