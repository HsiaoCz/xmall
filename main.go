package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/xmall/handlers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("./xmall.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	var (
		port         = os.Getenv("PORT")
		logger       = slog.New(slog.NewJSONHandler(io.MultiWriter(file, os.Stdout), &slog.HandlerOptions{}))
		userHandlers = &handlers.UserHandlers{}
		app          = http.NewServeMux()
	)
	// TODO: info error debug ervery log use itselves file
	slog.SetDefault(logger)

	{
		// router
		// user
		app.HandleFunc("POST /api/v1/user", handlers.TransferHandlerfunc(userHandlers.HandleCreateUser))

	}

	srv := http.Server{
		Addr:         port,
		Handler:      app,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	slog.Info("the server is running", "listen address", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			slog.Error("listen and serve error", "error message", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error", "error message", err)
		return
	}

	slog.Info("the server shutdown")

}
