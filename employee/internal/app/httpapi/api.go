package httpapi

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/BayMaxx2001/manager-employee/pkg/messaging/httppub"
	"github.com/BayMaxx2001/manager-employee/pkg/messaging/httpsub"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(mux *chi.Mux) {
	mux.Route("/api/v1", func(r chi.Router) {
		mux.Use(middleware.Timeout(30 * time.Second))
		mux.Use(middleware.Recoverer)

		r.Route("/employee", func(r chi.Router) {
			r.Get("/", GetAllEmployees)
			r.Post("/", CreateEmployee)
			r.Put("/{uid}", UpdateEmployee)
			r.Delete("/{uid}", DeleteEmployee)
			r.Get("/{uid}", FindEmployee)
		})
		r.Route("/event", func(r chi.Router) {
			r.Post("/{event}", AddEmployeeToTeam)
			r.Delete("/{event}", DeleteEmployeeToTeam)
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

	// pub-sub
	httppub.Init()
	httpsub.Init()
	publishAdd()
	publishDelete()

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

func publishAdd() { // oke
	url := url.URL{
		Scheme: "http",
		Host:   "localhost:8181",
		Path:   "/api/v1/event/employee-team",
	}
	pub := httppub.NewPublisher("employee-team", url, nil)
	httppub.ConnectPub(*pub, "employee-team")
}

func publishDelete() { // oke
	url := url.URL{
		Scheme: "http",
		Host:   "localhost:8181",
		Path:   "/api/v1/event/employee-team",
	}
	pub := httppub.NewPublisher("employee-team", url, nil)
	httppub.ConnectPub(*pub, "employee-team")
}
