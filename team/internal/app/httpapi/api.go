package httpapi

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

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