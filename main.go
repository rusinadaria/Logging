package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"os"
	"github.com/joho/godotenv"
)

func handlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HOME ROUT")
}

// func handlerGuid(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "GET GUID ROUT")
// }

// func handlerRefresh(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "REFRESH ROUT")
// }

func main() {
	godotenv.Load()
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
        slog.Error("Unable to open a file for writing")
    }

	opts := &slog.HandlerOptions{
        Level: slog.LevelDebug,
    }

	logger := slog.New(slog.NewJSONHandler(f, opts))
	logger.Info("Info message")

	ConnectDatabase(logger)

	router := chi.NewRouter()
    router.HandleFunc("/", handlerHome)
	router.Get("/auth/{guid}", handleGetGuidAndIP)
	router.Get("/auth/refresh", refreshHandler)

	// http.ListenAndServe(":80", nil)

	err = http.ListenAndServe(os.Getenv("PORT"), router)
	if err != nil {
		logger.Error("failed start server")
		panic(err)
	}

	// srv := &http.Server {
	// 	Addr: ":80",
	// 	Handler: router,
	// }

	// err = srv.ListenAndServe()
	// if err != nil {
	// 	fmt.Print("failed start server")
	// 	logger.Error("failed start server")
	// 	panic(err)
	// }
	
}