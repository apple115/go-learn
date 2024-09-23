package main

import (
	"context"
	"fmt"
	"go-gin-example/models"
	"go-gin-example/pkg/logging"
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
	setting.Setup()
	models.Setup()
	logging.Setup()
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
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
