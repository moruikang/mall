package main

import (
	"fmt"
	"log"
	"mall.com/config/global"
	"mall.com/pkg/initialize"
	"mall.com/routers"
	"net/http"
	"time"
)

func init() {

	initialize.LoadConfig()
	fmt.Println(global.Config.Mysql.Url)
	fmt.Println(global.Config.Mysql)
	initialize.Mysql()
	initialize.Redis()
	initialize.Elastic()

}

func main() {

	engine := routers.Router()

	readTimeout := 60 * time.Second
	writeTimeout := 60 * time.Second
	endPoint := fmt.Sprintf("0.0.0.0:%d", global.Config.Server.Post)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        engine,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
