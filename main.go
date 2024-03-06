package main

import (
	"blog/pkg/setting"
	"blog/routers"
	"fmt"
	"net/http"
)

func main() {

	router := routers.InitRouter()

	s := &http.Server{
		//Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Addr:           fmt.Sprintf(":%d", 8000),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
