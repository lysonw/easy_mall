package main

import (
	"context"
	v1 "easy_mall/api/v1"
	i "easy_mall/init"
	"easy_mall/script"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	i.InitMySQL()
	i.InitCache()
	script.ProductClickRecordDaemon()

	srv := &http.Server{
		Addr:    ":" + "6111",
		Handler: v1.NewRouter(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
	log.Println("server has been shut down")
}
