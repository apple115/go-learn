package main

import (
	"context"
	"fmt"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fvbock/endless"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	s := endless.NewServer(endPoint, routers.InitRouter())
	s.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen:%s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutdwon server...")
	ctx, concel := context.WithTimeout(context.Background(), 5*time.Second)
	defer concel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Sever exiting")
}
