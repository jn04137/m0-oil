package main

import (
	"log"
	"context"
	"net/http"
	"go.uber.org/fx"
	"github.com/go-chi/chi/v5"
)

// Beginning of handler stuff 
type Handler struct {
	// db *sql.DB
	// logger *log.Logger
}

func NewHandler() *Handler {
	log.Println("Creating Handler")
	return &Handler{}
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Home handler hit"))
}

func (h *Handler) SecondHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("second handler hit"))
}
// End of handler stuff

// Start of Router
func NewAppRouter(handler *Handler) chi.Router {
	log.Println("Creating a new Chi Router")

	r := chi.NewRouter()
	r.Get("/home", handler.HomeHandler)
	r.Get("/second", handler.SecondHandler)
	return r
}
// End of Router

// Start of server
func NewServer(lc fx.Lifecycle, router chi.Router) *http.Server {
	log.Println("In the New Server function")
	srv := &http.Server {
		Addr: ":8080",
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Starting http server on ", srv.Addr)
			go func() {
				log.Println("Listen and Serve HTTP Server")
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalln("HTTP server ListenAndServe err ", err)
				}
			}()
			return nil
		},
		OnStop: func (ctx context.Context) error {
			log.Println("Shutting down HTTP server...")
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
// End of Server

func main() {
	fx.New(
		fx.Provide(
			NewAppRouter,
			NewHandler,
			NewServer,
		),
		fx.Invoke(func(*http.Server){}),
	).Run()
}
