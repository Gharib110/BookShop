package application

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
)

var router = gin.Default()

func StartApplication() {
	sig := make(chan os.Signal)

	signal.Notify(sig, os.Interrupt)

	mapURLS()
	go func() {
		err := router.Run(":8080")
		if err != nil {
			log.Println(err.Error())
			return
		}
	}()

	<-sig
}
