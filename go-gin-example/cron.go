package main

import (
	"go-gin-example/models"
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting...")
	c := cron.New()

	c.AddFunc("******", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})

	c.AddFunc("******", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})
	c.Start()

	ti := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-ti.C:
			ti.Reset(time.Second * 10)
		}
	}
}
